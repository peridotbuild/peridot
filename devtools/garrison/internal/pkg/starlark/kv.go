package starlark

import "go.starlark.net/starlark"

type KV struct {
	Key   string
	Value starlark.Value
}

type KVs struct {
	KV       []KV
	indexLoc map[string]int
}

func NewKVs() *KVs {
	return &KVs{
		KV:       make([]KV, 0),
		indexLoc: make(map[string]int),
	}
}

func (k *KVs) Add(kv KV) {
	k.KV = append(k.KV, kv)
	k.indexLoc[kv.Key] = len(k.KV) - 1
}

func (k *KVs) Get(key string) (starlark.Value, bool) {
	loc, ok := k.indexLoc[key]
	if !ok {
		return nil, false
	}
	return k.KV[loc].Value, true
}

func (k *KVs) Set(key string, value starlark.Value) {
	loc, ok := k.indexLoc[key]
	if !ok {
		k.Add(KV{
			Key:   key,
			Value: value,
		})
		return
	}
	k.KV[loc].Value = value
}

func (k *KVs) Del(key string) {
	loc, ok := k.indexLoc[key]
	if !ok {
		return
	}
	k.KV = append(k.KV[:loc], k.KV[loc+1:]...)
	delete(k.indexLoc, key)
}

func kwargsToKV(kwargs []starlark.Tuple) KVs {
	kvs := NewKVs()
	for _, kv := range kwargs {
		kvs.Add(KV{
			Key:   kv.Index(0).(starlark.String).GoString(),
			Value: kv.Index(1),
		})
	}
	return *kvs
}
