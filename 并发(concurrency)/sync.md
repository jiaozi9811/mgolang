# sync

## sync.Map

type Map struct {
func (m *Map) Delete(key interface{}) {
func (m *Map) Load(key interface{}) (value interface{}, ok bool) {
func (m *Map) LoadOrStore(key, value interface{}) (actual interface{}, loaded bool) {
func (m *Map) Range(f func(key, value interface{}) bool) {          //循环读取map中的值
func (m *Map) Store(key, value interface{}) {