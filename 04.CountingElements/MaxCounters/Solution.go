package Ex04CoutingElementsMaxCounters

// https://codility.com/demo/results/trainingW8RGSB-J5M/

func Solution(N int, A []int) []int {
    counters := *NewCounters(N);
    
    for _, a := range A {
        counters.Apply(a);
    }
    
    return counters.GetCounters();
}

type counters struct {
    maxCounterValue int;
    currentMinCounterValue int;
    counters []int;
}

func NewCounters(size int) *counters {
    return &counters { counters: make([]int, size) };
}

func (c *counters) Apply(command int) {
    if command <= len(c.counters) {
        c.IncrementCounter(command - 1);
    } else {
        c.currentMinCounterValue = c.maxCounterValue;
    }
}

func (c *counters) GetCounter(index int) int {
    if (c.counters[index] <= c.currentMinCounterValue) {
        return c.currentMinCounterValue;
    }
    return c.counters[index];
}

func (c *counters) IncrementCounter(index int) {
    counterValue := c.GetCounter(index) + 1;
    c.counters[index] = counterValue;
    if (counterValue > c.maxCounterValue) {
        c.maxCounterValue = counterValue;
    }
}

func (c *counters) GetCounters() []int {
    result := make([]int, len(c.counters));
    for i := range c.counters {
        result[i] = c.GetCounter(i);
    }
    return result;
}