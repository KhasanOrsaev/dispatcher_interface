package dispatcher_interface

type Source interface {
	// ReadData получение данных из источника
	ReadData(inChannel chan <- map[interface{}][]byte) error
	// Confirm подтверждение о записи в назначение
	Confirm(confirmChannel <- chan interface{}) error
	// SetConfig инициализировать конфигурацию
	SetConfig(f []byte) error
	OpenConnection()error
	CloseConnection()error
}

type Destination interface {
	// WriteData запись данных в назначение
	WriteData(outChannel <- chan map[interface{}][]byte, confirmChannel chan <- interface{}, crashChannel chan <- []byte) error
	// SetConfig инициализировать конфигурацию
	SetConfig(f []byte) error
	OpenConnection()error
	CloseConnection()error
}

type Crash interface {
	// SaveData сохранение записи при которой произошла ошибка
	SaveData(crashChannel <- chan []byte)
	// SetConfig инициализировать конфигурацию
	SetConfig(f []byte) error
	OpenConnection()error
	CloseConnection()error
}