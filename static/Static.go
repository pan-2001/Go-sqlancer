package static

type Config struct {
	tableID         int
	columnID        int
	expressionDepth int
}

var c *Config

func Init() {
	c = new(Config)
	c.expressionDepth = 2
}

func (c *Config) WithConfig() *Config {
	return c
}

func WithConfig() *Config {
	return c.WithConfig()
}

func (c *Config) GetColumnID() int {
	return c.columnID
}

func (c *Config) PlusColumnID() {
	c.columnID++
}

func (c *Config) GetTableID() int {
	return c.tableID
}

func (c *Config) PlusTableID() {
	c.tableID++
}

func (c *Config) GetExpressionDepth() int {
	return c.expressionDepth
}
