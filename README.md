# Ayenda Services

## Atomic
When somethien it considered atomic or have this property, this means that within context that it
is operathing, it's indivisible or uninterruptible.

But en different context can be change a one example for this it's some operation in the context of
the operaty system can be atomic but in the application not and viceversa.

## Memmory Access Synchronization
It's means two concurrent proccess are attempting to access the same area of memory, and they accessing
in memory is not atomic. That's why we use sync.Mutex Let's see the next code

var memoryAccess sync.Mutex
var value int

go func() {
	memoryAccess.Lock() // Block procces for change value
	value++
	memoryAccess.Unlock() // unBlock procces after change value
}()
