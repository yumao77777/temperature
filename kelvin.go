package tempconv
import "fmt"
type cs float64
type kl float64

const (
    AbsoluteZeroK Celsius = -273.15
)

func CToK(c cs) kl { return kl(c + 273.15)}
func KToC(k kl) cs { return cs(k - 273.15)}


func (c cs) String() string    { return fmt.Sprintf("%g°C", c) }
func (k kl) String() string { return fmt.Sprintf("%g°F", k) }
