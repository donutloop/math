package calc

import (
	"fmt"
	"prototype_kl/calc/schema"
	"strings"
)

// schema.Topics maps function/constant/command names to one-line docs.

// help looks up a topic and prints its documentation.
func (c *Calculator) help(topic string) error {
	topic = strings.TrimSpace(topic)
	if topic == "" {
		c.printHelp()
		return nil
	}
	key := strings.ToLower(topic)
	if strings.HasPrefix(key, "@") {
		key = "@"
	}
	doc, ok := schema.Topics[key]
	if !ok {
		return fmt.Errorf("no help for %q; try 'help'", topic)
	}
	fmt.Fprintln(c.out, doc)
	return nil
}
