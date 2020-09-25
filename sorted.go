package sorted_set

import "sort"

// Uint32 start

type Uint32 []uint32

type implUint32Builder struct {
	valueSet map[uint32]bool
}

type Uint32Builder interface {
	Add(values ...uint32) Uint32Builder
	Sort() (out []uint32)
}

func NewUint32(values ...uint32) Uint32Builder {
	b := &implUint32Builder{
		valueSet: make(map[uint32]bool),
	}
	b.Add(values...)
	return b
}

func (b *implUint32Builder) Add(values ...uint32) Uint32Builder {
	for _, value := range values {
		b.valueSet[value] = true
	}
	return b
}

func (b *implUint32Builder) Sort() (out []uint32) {
	if b.valueSet == nil {
		return []uint32{}
	}
	out = make([]uint32, len(b.valueSet))
	valueIndex := 0
	for value := range b.valueSet {
		out[valueIndex] = value
		valueIndex++
	}
	sortable := sliceUint32(out)
	sort.Sort(&sortable)
	return
}

type sliceUint32 []uint32

func (v sliceUint32) Len() int {
	return len(v)
}

func (v sliceUint32) Less(i, j int) bool {
	return v[i] < v[j]
}

func (v *sliceUint32) Swap(i, j int) {
	(*v)[i], (*v)[j] = (*v)[j], (*v)[i]
}

// Uint32 end
// Uint8 start

type Uint8 []uint8

type implUint8Builder struct {
	valueSet map[uint8]bool
}

type Uint8Builder interface {
	Add(values ...uint8) Uint8Builder
	Sort() (out []uint8)
}

func NewUint8(values ...uint8) Uint8Builder {
	b := &implUint8Builder{
		valueSet: make(map[uint8]bool),
	}
	b.Add(values...)
	return b
}

func (b *implUint8Builder) Add(values ...uint8) Uint8Builder {
	for _, value := range values {
		b.valueSet[value] = true
	}
	return b
}

func (b *implUint8Builder) Sort() (out []uint8) {
	if b.valueSet == nil {
		return []uint8{}
	}
	out = make([]uint8, len(b.valueSet))
	valueIndex := 0
	for value := range b.valueSet {
		out[valueIndex] = value
		valueIndex++
	}
	sortable := sliceUint8(out)
	sort.Sort(&sortable)
	return
}

type sliceUint8 []uint8

func (v sliceUint8) Len() int {
	return len(v)
}

func (v sliceUint8) Less(i, j int) bool {
	return v[i] < v[j]
}

func (v *sliceUint8) Swap(i, j int) {
	(*v)[i], (*v)[j] = (*v)[j], (*v)[i]
}

// Uint8 end
// Float32 start

type Float32 []float32

type implFloat32Builder struct {
	valueSet map[float32]bool
}

type Float32Builder interface {
	Add(values ...float32) Float32Builder
	Sort() (out []float32)
}

func NewFloat32(values ...float32) Float32Builder {
	b := &implFloat32Builder{
		valueSet: make(map[float32]bool),
	}
	b.Add(values...)
	return b
}

func (b *implFloat32Builder) Add(values ...float32) Float32Builder {
	for _, value := range values {
		b.valueSet[value] = true
	}
	return b
}

func (b *implFloat32Builder) Sort() (out []float32) {
	if b.valueSet == nil {
		return []float32{}
	}
	out = make([]float32, len(b.valueSet))
	valueIndex := 0
	for value := range b.valueSet {
		out[valueIndex] = value
		valueIndex++
	}
	sortable := sliceFloat32(out)
	sort.Sort(&sortable)
	return
}

type sliceFloat32 []float32

func (v sliceFloat32) Len() int {
	return len(v)
}

func (v sliceFloat32) Less(i, j int) bool {
	return v[i] < v[j]
}

func (v *sliceFloat32) Swap(i, j int) {
	(*v)[i], (*v)[j] = (*v)[j], (*v)[i]
}

// Float32 end
// Uintptr start

type Uintptr []uintptr

type implUintptrBuilder struct {
	valueSet map[uintptr]bool
}

type UintptrBuilder interface {
	Add(values ...uintptr) UintptrBuilder
	Sort() (out []uintptr)
}

func NewUintptr(values ...uintptr) UintptrBuilder {
	b := &implUintptrBuilder{
		valueSet: make(map[uintptr]bool),
	}
	b.Add(values...)
	return b
}

func (b *implUintptrBuilder) Add(values ...uintptr) UintptrBuilder {
	for _, value := range values {
		b.valueSet[value] = true
	}
	return b
}

func (b *implUintptrBuilder) Sort() (out []uintptr) {
	if b.valueSet == nil {
		return []uintptr{}
	}
	out = make([]uintptr, len(b.valueSet))
	valueIndex := 0
	for value := range b.valueSet {
		out[valueIndex] = value
		valueIndex++
	}
	sortable := sliceUintptr(out)
	sort.Sort(&sortable)
	return
}

type sliceUintptr []uintptr

func (v sliceUintptr) Len() int {
	return len(v)
}

func (v sliceUintptr) Less(i, j int) bool {
	return v[i] < v[j]
}

func (v *sliceUintptr) Swap(i, j int) {
	(*v)[i], (*v)[j] = (*v)[j], (*v)[i]
}

// Uintptr end
// Uint start

type Uint []uint

type implUintBuilder struct {
	valueSet map[uint]bool
}

type UintBuilder interface {
	Add(values ...uint) UintBuilder
	Sort() (out []uint)
}

func NewUint(values ...uint) UintBuilder {
	b := &implUintBuilder{
		valueSet: make(map[uint]bool),
	}
	b.Add(values...)
	return b
}

func (b *implUintBuilder) Add(values ...uint) UintBuilder {
	for _, value := range values {
		b.valueSet[value] = true
	}
	return b
}

func (b *implUintBuilder) Sort() (out []uint) {
	if b.valueSet == nil {
		return []uint{}
	}
	out = make([]uint, len(b.valueSet))
	valueIndex := 0
	for value := range b.valueSet {
		out[valueIndex] = value
		valueIndex++
	}
	sortable := sliceUint(out)
	sort.Sort(&sortable)
	return
}

type sliceUint []uint

func (v sliceUint) Len() int {
	return len(v)
}

func (v sliceUint) Less(i, j int) bool {
	return v[i] < v[j]
}

func (v *sliceUint) Swap(i, j int) {
	(*v)[i], (*v)[j] = (*v)[j], (*v)[i]
}

// Uint end
// Uint64 start

type Uint64 []uint64

type implUint64Builder struct {
	valueSet map[uint64]bool
}

type Uint64Builder interface {
	Add(values ...uint64) Uint64Builder
	Sort() (out []uint64)
}

func NewUint64(values ...uint64) Uint64Builder {
	b := &implUint64Builder{
		valueSet: make(map[uint64]bool),
	}
	b.Add(values...)
	return b
}

func (b *implUint64Builder) Add(values ...uint64) Uint64Builder {
	for _, value := range values {
		b.valueSet[value] = true
	}
	return b
}

func (b *implUint64Builder) Sort() (out []uint64) {
	if b.valueSet == nil {
		return []uint64{}
	}
	out = make([]uint64, len(b.valueSet))
	valueIndex := 0
	for value := range b.valueSet {
		out[valueIndex] = value
		valueIndex++
	}
	sortable := sliceUint64(out)
	sort.Sort(&sortable)
	return
}

type sliceUint64 []uint64

func (v sliceUint64) Len() int {
	return len(v)
}

func (v sliceUint64) Less(i, j int) bool {
	return v[i] < v[j]
}

func (v *sliceUint64) Swap(i, j int) {
	(*v)[i], (*v)[j] = (*v)[j], (*v)[i]
}

// Uint64 end
// String start

type String []string

type implStringBuilder struct {
	valueSet map[string]bool
}

type StringBuilder interface {
	Add(values ...string) StringBuilder
	Sort() (out []string)
}

func NewString(values ...string) StringBuilder {
	b := &implStringBuilder{
		valueSet: make(map[string]bool),
	}
	b.Add(values...)
	return b
}

func (b *implStringBuilder) Add(values ...string) StringBuilder {
	for _, value := range values {
		b.valueSet[value] = true
	}
	return b
}

func (b *implStringBuilder) Sort() (out []string) {
	if b.valueSet == nil {
		return []string{}
	}
	out = make([]string, len(b.valueSet))
	valueIndex := 0
	for value := range b.valueSet {
		out[valueIndex] = value
		valueIndex++
	}
	sortable := sliceString(out)
	sort.Sort(&sortable)
	return
}

type sliceString []string

func (v sliceString) Len() int {
	return len(v)
}

func (v sliceString) Less(i, j int) bool {
	return v[i] < v[j]
}

func (v *sliceString) Swap(i, j int) {
	(*v)[i], (*v)[j] = (*v)[j], (*v)[i]
}

// String end
// Float64 start

type Float64 []float64

type implFloat64Builder struct {
	valueSet map[float64]bool
}

type Float64Builder interface {
	Add(values ...float64) Float64Builder
	Sort() (out []float64)
}

func NewFloat64(values ...float64) Float64Builder {
	b := &implFloat64Builder{
		valueSet: make(map[float64]bool),
	}
	b.Add(values...)
	return b
}

func (b *implFloat64Builder) Add(values ...float64) Float64Builder {
	for _, value := range values {
		b.valueSet[value] = true
	}
	return b
}

func (b *implFloat64Builder) Sort() (out []float64) {
	if b.valueSet == nil {
		return []float64{}
	}
	out = make([]float64, len(b.valueSet))
	valueIndex := 0
	for value := range b.valueSet {
		out[valueIndex] = value
		valueIndex++
	}
	sortable := sliceFloat64(out)
	sort.Sort(&sortable)
	return
}

type sliceFloat64 []float64

func (v sliceFloat64) Len() int {
	return len(v)
}

func (v sliceFloat64) Less(i, j int) bool {
	return v[i] < v[j]
}

func (v *sliceFloat64) Swap(i, j int) {
	(*v)[i], (*v)[j] = (*v)[j], (*v)[i]
}

// Float64 end
// Int start

type Int []int

type implIntBuilder struct {
	valueSet map[int]bool
}

type IntBuilder interface {
	Add(values ...int) IntBuilder
	Sort() (out []int)
}

func NewInt(values ...int) IntBuilder {
	b := &implIntBuilder{
		valueSet: make(map[int]bool),
	}
	b.Add(values...)
	return b
}

func (b *implIntBuilder) Add(values ...int) IntBuilder {
	for _, value := range values {
		b.valueSet[value] = true
	}
	return b
}

func (b *implIntBuilder) Sort() (out []int) {
	if b.valueSet == nil {
		return []int{}
	}
	out = make([]int, len(b.valueSet))
	valueIndex := 0
	for value := range b.valueSet {
		out[valueIndex] = value
		valueIndex++
	}
	sortable := sliceInt(out)
	sort.Sort(&sortable)
	return
}

type sliceInt []int

func (v sliceInt) Len() int {
	return len(v)
}

func (v sliceInt) Less(i, j int) bool {
	return v[i] < v[j]
}

func (v *sliceInt) Swap(i, j int) {
	(*v)[i], (*v)[j] = (*v)[j], (*v)[i]
}

// Int end
// Uint16 start

type Uint16 []uint16

type implUint16Builder struct {
	valueSet map[uint16]bool
}

type Uint16Builder interface {
	Add(values ...uint16) Uint16Builder
	Sort() (out []uint16)
}

func NewUint16(values ...uint16) Uint16Builder {
	b := &implUint16Builder{
		valueSet: make(map[uint16]bool),
	}
	b.Add(values...)
	return b
}

func (b *implUint16Builder) Add(values ...uint16) Uint16Builder {
	for _, value := range values {
		b.valueSet[value] = true
	}
	return b
}

func (b *implUint16Builder) Sort() (out []uint16) {
	if b.valueSet == nil {
		return []uint16{}
	}
	out = make([]uint16, len(b.valueSet))
	valueIndex := 0
	for value := range b.valueSet {
		out[valueIndex] = value
		valueIndex++
	}
	sortable := sliceUint16(out)
	sort.Sort(&sortable)
	return
}

type sliceUint16 []uint16

func (v sliceUint16) Len() int {
	return len(v)
}

func (v sliceUint16) Less(i, j int) bool {
	return v[i] < v[j]
}

func (v *sliceUint16) Swap(i, j int) {
	(*v)[i], (*v)[j] = (*v)[j], (*v)[i]
}

// Uint16 end
// Byte start

type Byte []byte

type implByteBuilder struct {
	valueSet map[byte]bool
}

type ByteBuilder interface {
	Add(values ...byte) ByteBuilder
	Sort() (out []byte)
}

func NewByte(values ...byte) ByteBuilder {
	b := &implByteBuilder{
		valueSet: make(map[byte]bool),
	}
	b.Add(values...)
	return b
}

func (b *implByteBuilder) Add(values ...byte) ByteBuilder {
	for _, value := range values {
		b.valueSet[value] = true
	}
	return b
}

func (b *implByteBuilder) Sort() (out []byte) {
	if b.valueSet == nil {
		return []byte{}
	}
	out = make([]byte, len(b.valueSet))
	valueIndex := 0
	for value := range b.valueSet {
		out[valueIndex] = value
		valueIndex++
	}
	sortable := sliceByte(out)
	sort.Sort(&sortable)
	return
}

type sliceByte []byte

func (v sliceByte) Len() int {
	return len(v)
}

func (v sliceByte) Less(i, j int) bool {
	return v[i] < v[j]
}

func (v *sliceByte) Swap(i, j int) {
	(*v)[i], (*v)[j] = (*v)[j], (*v)[i]
}

// Byte end
// Rune start

type Rune []rune

type implRuneBuilder struct {
	valueSet map[rune]bool
}

type RuneBuilder interface {
	Add(values ...rune) RuneBuilder
	Sort() (out []rune)
}

func NewRune(values ...rune) RuneBuilder {
	b := &implRuneBuilder{
		valueSet: make(map[rune]bool),
	}
	b.Add(values...)
	return b
}

func (b *implRuneBuilder) Add(values ...rune) RuneBuilder {
	for _, value := range values {
		b.valueSet[value] = true
	}
	return b
}

func (b *implRuneBuilder) Sort() (out []rune) {
	if b.valueSet == nil {
		return []rune{}
	}
	out = make([]rune, len(b.valueSet))
	valueIndex := 0
	for value := range b.valueSet {
		out[valueIndex] = value
		valueIndex++
	}
	sortable := sliceRune(out)
	sort.Sort(&sortable)
	return
}

type sliceRune []rune

func (v sliceRune) Len() int {
	return len(v)
}

func (v sliceRune) Less(i, j int) bool {
	return v[i] < v[j]
}

func (v *sliceRune) Swap(i, j int) {
	(*v)[i], (*v)[j] = (*v)[j], (*v)[i]
}

// Rune end
// Int64 start

type Int64 []int64

type implInt64Builder struct {
	valueSet map[int64]bool
}

type Int64Builder interface {
	Add(values ...int64) Int64Builder
	Sort() (out []int64)
}

func NewInt64(values ...int64) Int64Builder {
	b := &implInt64Builder{
		valueSet: make(map[int64]bool),
	}
	b.Add(values...)
	return b
}

func (b *implInt64Builder) Add(values ...int64) Int64Builder {
	for _, value := range values {
		b.valueSet[value] = true
	}
	return b
}

func (b *implInt64Builder) Sort() (out []int64) {
	if b.valueSet == nil {
		return []int64{}
	}
	out = make([]int64, len(b.valueSet))
	valueIndex := 0
	for value := range b.valueSet {
		out[valueIndex] = value
		valueIndex++
	}
	sortable := sliceInt64(out)
	sort.Sort(&sortable)
	return
}

type sliceInt64 []int64

func (v sliceInt64) Len() int {
	return len(v)
}

func (v sliceInt64) Less(i, j int) bool {
	return v[i] < v[j]
}

func (v *sliceInt64) Swap(i, j int) {
	(*v)[i], (*v)[j] = (*v)[j], (*v)[i]
}

// Int64 end
// Int32 start

type Int32 []int32

type implInt32Builder struct {
	valueSet map[int32]bool
}

type Int32Builder interface {
	Add(values ...int32) Int32Builder
	Sort() (out []int32)
}

func NewInt32(values ...int32) Int32Builder {
	b := &implInt32Builder{
		valueSet: make(map[int32]bool),
	}
	b.Add(values...)
	return b
}

func (b *implInt32Builder) Add(values ...int32) Int32Builder {
	for _, value := range values {
		b.valueSet[value] = true
	}
	return b
}

func (b *implInt32Builder) Sort() (out []int32) {
	if b.valueSet == nil {
		return []int32{}
	}
	out = make([]int32, len(b.valueSet))
	valueIndex := 0
	for value := range b.valueSet {
		out[valueIndex] = value
		valueIndex++
	}
	sortable := sliceInt32(out)
	sort.Sort(&sortable)
	return
}

type sliceInt32 []int32

func (v sliceInt32) Len() int {
	return len(v)
}

func (v sliceInt32) Less(i, j int) bool {
	return v[i] < v[j]
}

func (v *sliceInt32) Swap(i, j int) {
	(*v)[i], (*v)[j] = (*v)[j], (*v)[i]
}

// Int32 end
// Int16 start

type Int16 []int16

type implInt16Builder struct {
	valueSet map[int16]bool
}

type Int16Builder interface {
	Add(values ...int16) Int16Builder
	Sort() (out []int16)
}

func NewInt16(values ...int16) Int16Builder {
	b := &implInt16Builder{
		valueSet: make(map[int16]bool),
	}
	b.Add(values...)
	return b
}

func (b *implInt16Builder) Add(values ...int16) Int16Builder {
	for _, value := range values {
		b.valueSet[value] = true
	}
	return b
}

func (b *implInt16Builder) Sort() (out []int16) {
	if b.valueSet == nil {
		return []int16{}
	}
	out = make([]int16, len(b.valueSet))
	valueIndex := 0
	for value := range b.valueSet {
		out[valueIndex] = value
		valueIndex++
	}
	sortable := sliceInt16(out)
	sort.Sort(&sortable)
	return
}

type sliceInt16 []int16

func (v sliceInt16) Len() int {
	return len(v)
}

func (v sliceInt16) Less(i, j int) bool {
	return v[i] < v[j]
}

func (v *sliceInt16) Swap(i, j int) {
	(*v)[i], (*v)[j] = (*v)[j], (*v)[i]
}

// Int16 end
// Int8 start

type Int8 []int8

type implInt8Builder struct {
	valueSet map[int8]bool
}

type Int8Builder interface {
	Add(values ...int8) Int8Builder
	Sort() (out []int8)
}

func NewInt8(values ...int8) Int8Builder {
	b := &implInt8Builder{
		valueSet: make(map[int8]bool),
	}
	b.Add(values...)
	return b
}

func (b *implInt8Builder) Add(values ...int8) Int8Builder {
	for _, value := range values {
		b.valueSet[value] = true
	}
	return b
}

func (b *implInt8Builder) Sort() (out []int8) {
	if b.valueSet == nil {
		return []int8{}
	}
	out = make([]int8, len(b.valueSet))
	valueIndex := 0
	for value := range b.valueSet {
		out[valueIndex] = value
		valueIndex++
	}
	sortable := sliceInt8(out)
	sort.Sort(&sortable)
	return
}

type sliceInt8 []int8

func (v sliceInt8) Len() int {
	return len(v)
}

func (v sliceInt8) Less(i, j int) bool {
	return v[i] < v[j]
}

func (v *sliceInt8) Swap(i, j int) {
	(*v)[i], (*v)[j] = (*v)[j], (*v)[i]
}

// Int8 end
