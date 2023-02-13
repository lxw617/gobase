package reentrant_mutex

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

const KEY = "EXAMPLE_LOCK"

// Lock 用于测试的锁
type Lock struct {
	// redis连接池
	Rdb *redis.Client
	// hash锁key
	Key string
	// hash锁field(随机数,实时唯一)
	Field int
	// 锁有效期
	Expiration time.Duration
	// 用于测试的初始递归层数
	RecursionLevel int
	// 用于测试的最大递归层数
	MaxRecursionLevel int
	// 用于测试的任务最小执行时间
	Min int
	// 用于测试的任务最大执行时间
	Max int
	// 加锁失败的重试间隔
	RetryInterval time.Duration
	// 加锁失败的重试次数
	RetryTimes int
	// 继承*sync.Once的特性
	*sync.Once
	// 用于测试打印的线程标签
	Tag string
}

func init() {
	fmt.Println("initializing rand seed for rand testing...")
	rand.Seed(time.Now().UnixNano())
}

// 生成一个随机标签
func getRandTag(n int) string {
	var runes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	tag := make([]rune, n)
	for i := range tag {
		tag[i] = runes[rand.Intn(len(runes))]
	}
	return string(tag)
}

// NewLock 初始化
func NewLock(rdb *redis.Client) *Lock {
	l := Lock{
		Rdb:               rdb,
		Key:               KEY, // 固定值
		Field:             rand.Int(),
		Expiration:        time.Millisecond * 200,
		RecursionLevel:    1,
		MaxRecursionLevel: 1,
		Min:               50,
		Max:               100,
		RetryInterval:     time.Millisecond * 50,
		RetryTimes:        5,
		Once:              new(sync.Once),
		Tag:               getRandTag(2),
	}
	return &l
}

// MockBusiness 模拟分布式业务加锁场景
func (l *Lock) MockBusiness() {
	fmt.Printf("%s的第%d次调用,Field:%d\n", l.Tag, l.RecursionLevel, l.Field)

	// 初始化仅用于当前调用的ctx,避免在重入调用完成后执行cancel()导致的上层调用出现context canceled错误
	var ctx, cancel = context.WithCancel(context.Background())

	defer func() {
		// 延迟停止守护线程
		cancel()
	}()

	set, err := l.lock(ctx)

	if err != nil {
		fmt.Println(l.Tag + " 加锁失败:" + err.Error())
		return
	}

	// 加锁失败,重试
	if set == false {
		res, err := l.retry(ctx)
		if err != nil {
			fmt.Println(l.Tag + " 重试加锁失败:" + err.Error())
			return
		}
		// 重试达到最大次数
		if res == false {
			fmt.Println(l.Tag + " server unavailable, try again later")
			return
		}
	}

	fmt.Println(l.Tag + "成功加锁")

	// 加锁成功,通过守护线程自动续期(此处可以异步执行,即使自动续期还没来得及执行业务就已经完成,也不会影响流程)
	go l.watchDog(ctx)

	fmt.Println(l.Tag + "等待业务处理完成...")
	// 模拟处理业务(通过随机时间模拟业务延迟)
	time.Sleep(time.Duration(rand.Intn(l.Max-l.Min)+l.Min) * time.Millisecond)

	// 模拟重入调用(测试锁的可重入)
	if l.RecursionLevel <= l.MaxRecursionLevel {
		l.RecursionLevel += 1
		l.MockBusiness()
	}

	// 业务处理完成
	// 释放锁
	val, err := l.unlock(ctx)
	if err != nil {
		fmt.Println(l.Tag + "锁释放失败:" + err.Error())
		return
	}

	// 递归调用中的结果都是false,因为lua脚本中的if分支counter>0,没有释放
	fmt.Println(l.Tag+"释放结果:", val)
}

// 守护线程(通过sync.Once.Do确保仅在线程第一次调用时执行自动续期)
func (l *Lock) watchDog(ctx context.Context) {
	l.Once.Do(func() {
		fmt.Printf("打开了%s的守护线程\n", l.Tag)
		for {
			select {
			// 业务完成
			case <-ctx.Done():
				fmt.Printf("%s任务完成,关闭%s的自动续期\n", l.Tag, l.Key)
				return
				// 业务未完成
			default:
				// 自动续期
				l.Rdb.PExpire(ctx, l.Key, l.Expiration)
				// 继续等待
				time.Sleep(l.Expiration / 2)
			}
		}
	})
}

// 加锁
func (l *Lock) lock(ctx context.Context) (res bool, err error) {
	lua := `
-- KEYS[1]:锁对应的key
-- ARGV[1]:锁的expire
-- ARGV[2]:锁对应的计数器field(随机值,防止误解锁),记录当前线程已加锁的次数
-- 判断锁是否空闲
if (redis.call('EXISTS', KEYS[1]) == 0) then
    -- 线程首次加锁(锁的初始化,值和过期时间)
    redis.call('HINCRBY', KEYS[1], ARGV[2], 1);
    redis.call('PEXPIRE', KEYS[1], ARGV[1]);
    return 1;
end;
-- 判断当前线程是否持有锁(锁被某个线程持有,通常是程序第N次(N>1)在线程内调用时会执行到此处)
if (redis.call('HEXISTS', KEYS[1], ARGV[2]) == 1) then
    -- 调用次数递增
    redis.call('HINCRBY', KEYS[1], ARGV[2], 1);
    -- 不处理续期,通过守护线程续期
    return 1;
end;
-- 锁被其他线程占用,加锁失败
return 0;
`

	scriptKeys := []string{l.Key}

	val, err := l.Rdb.Eval(ctx, lua, scriptKeys, int(l.Expiration), l.Field).Result()
	if err != nil {
		return
	}

	res = val == int64(1)

	return
}

// 解锁
func (l *Lock) unlock(ctx context.Context) (res bool, err error) {
	lua := `
-- KEYS[1]:锁对应的key
-- ARGV[1]:锁对应的计数器field(随机值,防止误解锁),记录当前线程已加锁的次数
-- 判断 hash set 是否存在
if (redis.call('HEXISTS', KEYS[1], ARGV[1]) == 0) then
    -- err = redis.Nil
    return nil;
end;
-- 计算当前可重入次数
local counter = redis.call('HINCRBY', KEYS[1], ARGV[1], -1);
if (counter > 0) then
-- 同一线程内部多次调用完成后尝试释放锁会进入此if分支
    return 0;
else
-- 同一线程最外层(第一次)调用完成后尝试释放锁会进入此if分支
-- 小于等于 0 代表内层嵌套调用已全部完成，可以解锁
    redis.call('DEL', KEYS[1]);
    return 1;
end;
-- err = redis.Nil
return nil;
`

	scriptKeys := []string{l.Key}
	val, err := l.Rdb.Eval(ctx, lua, scriptKeys, l.Field).Result()
	if err != nil {
		return
	}

	res = val == int64(1)

	return
}

// 重试
func (l *Lock) retry(ctx context.Context) (res bool, err error) {
	i := 1
	for i <= l.RetryTimes {
		fmt.Printf(l.Tag+"第%d次重试加锁中,Field:%d\n", i, l.Field)
		res, err = l.lock(ctx)

		if err != nil {
			return
		}

		if res == true {
			return
		}

		time.Sleep(l.RetryInterval)
		i++
	}
	return
}
