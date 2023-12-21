1. 一个线程安全集合，内部通过原子访问和锁机制实现结合的线程安全
   
# sync.map 应用场景
1. 适合读多写少的应用场景
2. 在key值以存在的情况下，可以无锁修改其value，比普通 map + 锁性能更好
   
# sync.Pool 作用 
1. 创建一个临时对象池，缓存一组对象用于重复利用，以此来减少内存分配和降低GC的动压力
   
# sync.Pool 应用场景
1. 可用于连接池（grpc 客户端、网络连接等）等场景
   
# sync.Pool 注意事项
1. 用于缓存一些创建成本较高，使用比较频繁的对象
2. Pool的长度默认为机器CPU线程数
3. 存储在Pool中的对象随时都可能在不被通知的情况下被回收
4. 没有什么创建成本的对象不建议使用对象池
   
# sync.Once 作用
1. 用来初始化单例资源

# sync.Once 使用场景
1. 单例场景
2. 仅加载一次的数据懒加载场景