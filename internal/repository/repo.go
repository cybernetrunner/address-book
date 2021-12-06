package repository

import (
	"github.com/cyneruxyz/address-book/internal/dummy"
	"hash/fnv"
	"strings"
	"sync"
)

type AddressBook struct {
	sync.Mutex
	Names 	map[string][]*dummy.AddressField
	Addrs 	map[string][]*dummy.AddressField
	Phones 	map[string][]*dummy.AddressField
	Hashes 	map[uint32]*dummy.AddressField
}

func NewAddressBook() *AddressBook {
	return &AddressBook{}
}

func (ab *AddressBook) CreateAddressField(af dummy.AddressField) (ok bool) {
	hash := getHash(af)
	if _, ok := ab.Hashes[hash]; ok {
		return !ok
	}

	ab.Lock()
	defer ab.Unlock()

	ab.Names[af.Name]	 = append(ab.Names[af.Name], &af)
	ab.Addrs[af.Address] = append(ab.Addrs[af.Address], &af)
	ab.Phones[af.Phone]  = append(ab.Phones[af.Phone], &af)
	ab.Hashes[hash] 	 = &af


	return true
}

func (ab *AddressBook) GetAddressFields(param string) (fields []*dummy.AddressField) {
	ab.Lock()
	defer ab.Unlock()

	if res, ok := ab.Names[param]; ok {
		return res
	} else if res, ok := ab.Addrs[param]; ok {
		return res
	} else if res, ok := ab.Phones[param]; ok {
		return res
	}

	return nil
}

func (ab *AddressBook) UpdateAddressField(original, replace dummy.AddressField) (ok bool) {
	hash := getHash(original)
	if _, ok := ab.Hashes[hash]; !ok {
		return !ok
	}
	delete(ab.Hashes, hash)

	ab.Lock()
	defer ab.Unlock()

	ab.Hashes[getHash(replace)] = &replace
	replaceField(ab.Names[original.Name], &original, &replace)
	replaceField(ab.Addrs[original.Address], &original, &replace)
	replaceField(ab.Phones[original.Phone], &original, &replace)

	return true

}

func (ab *AddressBook) DeleteAddressField(af dummy.AddressField) (ok bool) {
	hash := getHash(af)
	if _, ok := ab.Hashes[hash]; !ok {
		return !ok
	}

	ab.Lock()
	defer ab.Unlock()

	delete(ab.Hashes, hash)
	deleteField(ab.Names[af.Name], &af)
	deleteField(ab.Addrs[af.Address], &af)
	deleteField(ab.Phones[af.Phone], &af)


	return true
}

func replaceField(arr []*dummy.AddressField, o, r *dummy.AddressField){
	for i, v := range arr {
		if v == o {
			arr[i] = r
		}
	}
}

func deleteField(arr []*dummy.AddressField, f *dummy.AddressField){
	for i, v := range arr {
		if v == f {
			arr = append(arr[:i], arr[i+1:]...)
		}
	}
}

func getHash(af dummy.AddressField) uint32 {
	var sb strings.Builder

	sb.WriteString(af.Name)
	sb.WriteString(af.Address)
	sb.WriteString(af.Phone)

	h := fnv.New32a()
	_, _ = h.Write([]byte(sb.String()))

	return h.Sum32()
}
