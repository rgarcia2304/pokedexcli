package persist 

import "sync"

var lock sync.Mutex

var Marshal = func(v interface{}) (io.Reader, error){
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil{
		return nil, err
	}
	return bytes.NewReader(b), nil
}
//Save saves a representation of v to the file at path. 

func Save(path string, v interface{}) error{
	lock.Lock()
	defer lock.Unlock()
	f, err := os.Create(path)
	if err != nil{
		return err
	}
	defer f.Close()
	r, err := Marshal(v)
	if err != nil{
		return err
	}

	_, err = io.Copy(f, r)
	return err
}

func Load(path string, v interface{}) error{
	lock.Lock()
	defer lock.Unlock()
	f, err := os.Open(path)
	if err != nil{
		return err
	}
	defer f.Close()
	return Unmarshal(f,v)
}




