# new和make的区别

[TCO]

## 变量的声明

变量的声明我们可以通过var关键字，然后就可以在程序中使用。当我们不指定变量的默认值时，这些变量的默认值都是它们的零值，比如int类型的零值是0，string类型的零值是""，引用类型的零值是nil

## new

func new(Type) *Type
new函数接受一个参数，此参数是一个类型，分配好内存后，返回一个指向该类型内存地址的指针。同时把分配的内存设置为零

## make

func make(t Type,size ...IntegerType) Type

make也是用于内存分配的，但是和new不同，它只用于chan、map以及切片的内存创建，而且它返回的类型就是这三个类型本身，而不是它们的指针，因为这三种类型就是引用类型，所以就没必要返回它们的指针了。
注意，因为这三种类型是引用类型，所以必须得初始化，但是不是置为零值，这个和new是不一样的

## 二者异同

二者都是内存的分配（堆上），但是make只用于slice、map以及channel的初始化（非零值）；而new用于类型的内存分配，并且内存置为零。
make返回的还是这三个引用类型本身；而new返回的是指向类型的指针