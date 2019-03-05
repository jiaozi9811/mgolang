# sync/atomic

[TOC]

提供了底层的原子级内存操作，对于同步算法的实现很有用

### type Value

func (v *Value) Load() (x interface{})
func (v *Value) Store(x interface{})

atomic提供了五类原子操作:

    - Add, 增加和减少
    - CompareAndSwap, 比较并交换
    - Swap, 交换
    - Load , 读取
    - Store, 存储

## 原子增值

//对变量值进行原子增操作，并返回增加后的值
atomic.AddUint32(addr *uint32, delta uint32) uint32
atomic.AddUint64(addr *uint64, delta uint64) uint64
atomic.AddInt32(addr *int32, delta int32) int32
atomic.AddInt64(addr *int64, delta int64) int64
atomic.AddUintptr(addr *uintptr, delta uintptr) uintptr

## Compare And Swap(比较并交换)

//比较变量的值是否等于给定旧值，等于旧值的情况下才赋予新值，最后返回新值是否设置成功
atomic.CompareAndSwapUint32(addr *uint32, old, new uint32) bool
atomic.CompareAndSwapUint64(addr *uint64, old, new uint64) bool
atomic.CompareAndSwapInt32(addr *int32, old, new int32) bool
atomic.CompareAndSwapInt64(addr *int64, old, new int64) bool
atomic.CompareAndSwapUintptr(addr *uintptr, old, new uintptr) bool
atomic.CompareAndSwapPointer(addr *unsafe.Pointer, old, new unsafe.Pointer)

## 导出值

//导出变量当前的值
     在进行读取value的操作的过程中,其他对此值的读写操作是可以被同时进行的,那么这个读操作很可能会读取到一个只被修改了一半的数据
    Load函数保证了数据的原子性
atomic.LoadUint32(addr *uint32) uint32
atomic.LoadUint64(addr *uint64) uint64
atomic.LoadInt32(addr *int32) int32
atomic.LoadInt64(addr *int64) int64
atomic.LoadUintptr(addr *uintptr) uintptr
atomic.LoadPointer(addr *unsafe.Pointer) unsafe.Pointer

## 导入值

//赋予变量新值，而不管它原来是什么值
atomic.StoreUint32(addr *uint32, val uint32)
atomic.StoreUint64(addr *uint64, val uint64)
atomic.StoreInt32(addr *int32, val int32)
atomic.StoreInt64(addr *int64, val int64)
atomic.StoreUintptr(addr *uintptr, val uintptr)
atomic.StorePointer(addr *unsafe.Pointer, val unsafe.Pointer)

## 交换值

//赋予变量新值，同时返回变量的旧值
atomic.SwapUint32(addr *uint32, new uint32) old uint32
atomic.SwapUint64(addr *uint64, new uint64) old uint64
atomic.SwapInt32(addr *int32, new int32) old int32
atomic.SwapInt64(addr *int64, new int64) old int64
atomic.SwapUintptr(addr *uintptr, new uintptr) old uintptr
atomic.SwapPointer(addr *unsafe.Pointer, new unsafe.Pointer) old unsafe.Pointer

## int各种类型取值范围

|类型|长度(字节)|值范围|
|---|---|---|
|int8|1|-128 ~ 127|
|uint8(byte)|1|0 ~ 255
int16|2|32768~32767
uint16|2|0~65535
int32|4|2147483648~2147483647
uint32|4|0~4294967295
int64|8|-9223372036854775808~9223372036854775807
uint64|8|0~18446744073709551615
int|平台相关|平台相关
uint|平台相关|平台相关
uintptr|同指针|在32位平 下为4字节,64位平 下为8字节
