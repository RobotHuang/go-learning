package config

type AppConfig struct {
	KafkaConfig `ini:"kafka"`
	EtcdConfig `ini:"etcd"`
}

type KafkaConfig struct  {
	Address string `ini:"address"`
	//Topic string `ini:"topic"`
	ChanMaxSize int `ini:"chan_max_size"`
}

type TailLogConfig struct {
	FileName string `ini:"filename"`
}

type EtcdConfig struct {
	 Address string `ini:"address"`
	 Timeout int `ini:"timeout"`
	 Key string `ini:"collect_log_key"`
}
