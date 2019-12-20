package main

type config struct {
	Server *server `toml:server`
	ICP    []*icp  `toml:icp`
}

type server struct {
	Port string `toml:port`
}

type icp struct {
	URL string `toml:url`
	No  string `toml:no`
}
