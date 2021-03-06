type Mutex(互斥锁)
type Mutex struct {
func (m *Mutex) Lock()
func (m *Mutex) Unlock()


type RWMutex
type RWMutex struct {
写入锁
func (rw *RWMutex) Lock()
func (rw *RWMutex) Unlock()

读取锁
func (rw *RWMutex) RLock()
func (rw *RWMutex) RUnlock()

写解锁会唤醒所有因欲进行读锁定而被阻塞的goroutine。
读解锁只会在已无任何读锁定的情况下试图唤醒一个因欲进行写锁定而被阻塞的goroutine
若对一个未被写锁定的读写锁进行写解锁，会引发一个不可恢复的运行时恐慌，而对一个未被读锁定的读写锁进行读解锁同样会引起恐慌



sync.Once 执行一次
type Once struct {    // 包含隐藏或非导出字段 }

func (o *Once) Do(f func())
sync.Once类型的典型应用场景是执行仅需执行一次的任务
有些任务不适合在init函数中执行，此时sync.Once就派上用场了



sync.WaitGroup
type WaitGroup struct {  
func (wg *WaitGroup) Add(delta int)
func (wg *WaitGroup) Done()   //减少计数器的值
func (wg *WaitGroup) Wait()   //计数器为0，此方法才会执行

Add方法向内部计数器加delta，delta可为负数
如内部计数器为0，Wait方法将阻塞的所有线程释放。
Wait方法会检查计数器，如果计算为0，改方法返回，对后续程序的运行无影响；如计数器大于0，改方法所在的goroutine就会阻塞
如计数器小于0，报panic
Add方法需在Wait之前
sync.Add(-1)与sync.Done()相同



sync.Pool(临时对象池)
type Pool struct {    New func() interface{}    }

func (p *Pool) Get() interface{}
func (p *Pool) Put(x interface{})

Pool用于存储临时对象，将使用完毕的对象存入对象池，需要时取出重复使用，以避免重复创建相同的对象造成GC负担过重
存储的对象可以使用GC回收(如果该对象不再被引用)

如Pool中无数据，取数据是返回nil



sync.Coud(条件等待)
type Cond struct {
  L Locker
}
func NewCond(l Locker) *Cond  // 创建一个条件等待
func (c *Cond) Broadcast()
func (c *Cond) Signal()
func (c *Cond) Wait()

## sync.Map

type Map struct {
func (m *Map) Delete(key interface{}) {
func (m *Map) Load(key interface{}) (value interface{}, ok bool) {
func (m *Map) LoadOrStore(key, value interface{}) (actual interface{}, loaded bool) {
func (m *Map) Range(f func(key, value interface{}) bool) {          //循环读取map中的值
func (m *Map) Store(key, value interface{}) {