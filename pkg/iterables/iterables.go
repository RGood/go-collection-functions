package iterables

type SliceIterable[K interface{}] struct {
	abstractIterable[K]
	data  []K
	index int
}

func NewSliceIterable[K interface{}](elements []K) Iterable[K] {
	iter := &SliceIterable[K]{data: elements, index: 0}
	iter.Iterable = iter
	return iter
}

func (si *SliceIterable[K]) Next() (K, bool) {
	if si.index < len(si.data) {
		val := si.data[si.index]
		si.index++
		return val, true
	}

	var i K
	return i, false
}

func (si *SliceIterable[K]) Reset() {
	si.index = 0
}

type Pair[K comparable, V interface{}] struct {
	Key   K
	Value V
}

type MapIterable[K comparable, V interface{}] struct {
	abstractIterable[*Pair[K, V]]
	data          map[K]V
	index         int
	generatorChan chan *Pair[K, V]
}

func NewMapIterable[K comparable, V interface{}](data map[K]V) Iterable[*Pair[K, V]] {
	iter := &MapIterable[K, V]{
		data:          data,
		index:         0,
		generatorChan: make(chan *Pair[K, V]),
	}

	iter.Reset()

	iter.Iterable = iter
	return iter
}

func (mi *MapIterable[K, V]) Next() (*Pair[K, V], bool) {
	res, ok := <-mi.generatorChan
	return res, ok
}

func (mi *MapIterable[K, V]) Reset() {
	mi.generatorChan = make(chan *Pair[K, V])
	go func() {
		for k, v := range mi.data {
			mi.generatorChan <- &Pair[K, V]{Key: k, Value: v}
		}
		close(mi.generatorChan)
	}()
}

type MappingIterable[K, V interface{}] struct {
	abstractIterable[K]
	parentIterable Iterable[V]
	mapper         func(V) K
}

func (mi *MappingIterable[K, V]) Next() (K, bool) {
	element, ok := mi.parentIterable.Next()
	if !ok {
		var i K
		return i, false
	}

	return mi.mapper(element), true
}

func (mi *MappingIterable[K, V]) Reset() {
	mi.parentIterable.Reset()
}

func Map[K, V interface{}](parent Iterable[K], mapper func(K) V) Iterable[V] {
	iter := &MappingIterable[V, K]{
		parentIterable: parent,
		mapper:         mapper,
	}

	iter.Iterable = iter
	return iter
}

type FilterIterable[K interface{}] struct {
	abstractIterable[K]
	parentIterable Iterable[K]
	filter         func(K) bool
}

func (fi *FilterIterable[K]) Next() (K, bool) {
	element, ok := fi.parentIterable.Next()
	for ok {
		if fi.filter(element) {
			return element, true
		}

		element, ok = fi.parentIterable.Next()
	}

	var i K
	return i, false
}

func (fi *FilterIterable[K]) Reset() {
	fi.parentIterable.Reset()
}

func Filter[K interface{}](parent Iterable[K], filter func(K) bool) Iterable[K] {
	iter := &FilterIterable[K]{
		parentIterable: parent,
		filter:         filter,
	}

	iter.Iterable = iter
	return iter
}

func Reduce[K, V interface{}](elements Iterable[V], init K, reducer func(K, V) K) K {
	result := init
	elements.ForEach(func(_ int, element V) {
		result = reducer(result, element)
	})

	return result
}

type abstractIterable[K interface{}] struct {
	Iterable[K]
}

func (ai *abstractIterable[K]) ForEach(callback func(int, K)) {
	index := 0
	element, ok := ai.Next()
	for ok {
		callback(index, element)
		element, ok = ai.Next()
		index++
	}
}

func (ai *abstractIterable[K]) Last() (K, bool) {
	var e K
	element, ok := ai.Next()
	for ok {
		e = element
		element, ok = ai.Next()
	}

	return e, ok
}

func (ai *abstractIterable[K]) Collect() []K {
	res := []K{}
	ai.ForEach(func(_ int, element K) {
		res = append(res, element)
	})

	return res
}

type Iterable[K interface{}] interface {
	Next() (K, bool)
	ForEach(func(int, K))
	Last() (K, bool)
	Reset()
}
