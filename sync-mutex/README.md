1. 应用于一发多收的场景，即一组协程需要等待某一个协程完成一些前置准备的情况

# sync.Cond 注意事项   
1. 被叫方必须持有锁
2. 主叫方可以持有锁，但允许不持有
3. 尽可能 的减少无效的唤醒

# Mutex 与 RWMutex 作用
1. 并发场景下，通过锁机制，解决数据竟争的问题
   
# Mutex 与 RWMutex 应用场景
1. 协程安全
2. 数据竟争

# 注意事项 
1. 尽量避免使用锁
2. 应合理使用锁机制，不要滥用
   
# sync.Map
1.一个线程安全集合，内部通过原子访问和锁机制实现结合的线程安全

# sync.map 应用场景
1. 适合读多写少的应用场景
2. 在key值以存在的情况下，可以无所修改其 value，比 普通map + 锁性能更好


