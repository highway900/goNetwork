package goNetwork

type OrderedMap struct {
    m map[interface{}]interface{}
    order []interface{}
}

func (o *OrderedMap) Set(key interface{}, val interface{}) {
    o.order = append(o.order, key)
    o.m[key] = val
}

func (o *OrderedMap) Iter() <-chan struct {Key, Val interface{}} {
    ch := make(chan struct {Key, Val interface{}})
    go func() {
        for _, key := range o.order {
            ch <- struct {Key, Val interface{}} {key, o.m[key]}
        }
        close(ch)
    }()
    return ch
}
