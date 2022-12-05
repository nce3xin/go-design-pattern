package decorator

// InputStream 接口
type InputStream interface {
	Read([]byte) int
	Close()
}

// FileInputStream InputStream接口的实现类
type FileInputStream struct {
}

func (f FileInputStream) Read(bytes []byte) int {
	return 0
}

func (f FileInputStream) Close() {
	return
}

// BufferedInputStream 装饰类，实现基于缓存的读数据接口
type BufferedInputStream struct {
	in InputStream
}

func (b BufferedInputStream) Read(bytes []byte) int {
	//TODO 实现增强功能：基于缓存的读数据接口
	b.Read(bytes)
	//TODO 实现增强功能：基于缓存的读数据接口
	return 0
}

func (b BufferedInputStream) Close() {
	//TODO 实现增强功能：基于缓存的读数据接口
	b.Close()
	//TODO 实现增强功能：基于缓存的读数据接口
}

// DataInputStream 装饰类，实现读取基本类型数据的接口
type DataInputStream struct {
	in InputStream
}

func (d DataInputStream) Read(bytes []byte) int {
	//TODO 实现增强功能：基于读取基本类型数据的接口
	d.Read(bytes)
	//TODO 实现增强功能：基于读取基本类型数据的接口
	return 0
}

func (d DataInputStream) Close() {
	//TODO 实现增强功能：基于读取基本类型数据的接口
	d.Close()
	//TODO 实现增强功能：基于读取基本类型数据的接口
}

func (d DataInputStream) ReadInt() {
	return
}
