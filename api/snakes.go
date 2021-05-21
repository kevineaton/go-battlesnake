package api

// SnakeOptions are the different personalization options for a snake: https://play.battlesnake.com/references/customizations/
type SnakeOptions struct {
	Color string
	Head  string
	Tail  string
}

// Battlesnake is a single instance of a snake: https://docs.battlesnake.com/references/api#battlesnake
type Battlesnake struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Health  int     `json:"health"`
	Body    []Point `json:"body"`
	Latency string  `json:"latency"`
	Head    Point   `json:"head"`
	Length  int     `json:"length"`
	Shout   string  `json:"shout"`
	Squad   string  `json:"squad"`
}

// we specify all of the snake options here; the constants should be of SNAKE_WHAT_GROUP_NAME format. Only public customizations
// available for everyone are added here; feel free to override. Accurate as of 20210518: https://play.battlesnake.com/references/customizations/
const (
	// TODO: add all of the rest
	// heads
	SNAKE_HEAD_STANDARD_DEFAULT   = "default"
	SNAKE_HEAD_STANDARD_BELUGA    = "beluga"
	SNAKE_HEAD_STANDARD_BENDR     = "bendr"
	SNAKE_HEAD_STANDARD_DEAD      = "dead"
	SNAKE_HEAD_STANDARD_EVIL      = "evil"
	SNAKE_HEAD_STANDARD_FANG      = "fang"
	SNAKE_HEAD_STANDARD_PIXEL     = "pixel"
	SNAKE_HEAD_STANDARD_SAFE      = "safe"
	SNAKE_HEAD_STANDARD_SAND_WORM = "sand-worm"
	SNAKE_HEAD_STANDARD_SHADES    = "shades"
	SNAKE_HEAD_STANDARD_SILLY     = "silly"
	SNAKE_HEAD_STANDARD_SMILE     = "tongue"

	// heads - winter 2019
	SNAKE_HEAD_WINTER_2019_BONHOMME  = "bonhomme"
	SNAKE_HEAD_WINTER_2019_EARMUFFS  = "earmuffs"
	SNAKE_HEAD_WINTER_2019_RUDOLPH   = "rudolph"
	SNAKE_HEAD_WINTER_2019_SCARF     = "scarf"
	SNAKE_HEAD_WINTER_2019_SKI       = "ski"
	SNAKE_HEAD_WINTER_2019_SNOWMAN   = "snowman"
	SNAKE_HEAD_WINTER_2019_SNOW_WORM = "snow-worm"

	// heads - stay home and code 2020
	SNAKE_HEAD_CODE_2020_CAFFEINE   = "caffeine"
	SNAKE_HEAD_CODE_2020_GAMER      = "gamer"
	SNAKE_HEAD_CODE_2020_TIGER_KING = "tiger-king"
	SNAKE_HEAD_CODE_2020_WORKOUT    = "workout"

	// tails

	SNAKE_TAIL_STANDARD_DEFAULT      = "default"
	SNAKE_TAIL_STANDARD_BLOCK_BUM    = "block-bum"
	SNAKE_TAIL_STANDARD_BOLT         = "bolt"
	SNAKE_TAIL_STANDARD_CURLED       = "curled"
	SNAKE_TAIL_STANDARD_FAT_RATTLE   = "fat-rattle"
	SNAKE_TAIL_STANDARD_FRECKLED     = "freckled"
	SNAKE_TAIL_STANDARD_HOOK         = "hook"
	SNAKE_TAIL_STANDARD_PIXEL        = "pixel"
	SNAKE_TAIL_STANDARD_ROUND_BUM    = "round-bum"
	SNAKE_TAIL_STANDARD_SHARP        = "sharp"
	SNAKE_TAIL_STANDARD_SKINNY       = "skinny"
	SNAKE_TAIL_STANDARD_SMALL_RATTLE = "small-rattle"

	// tails - winter 2019
	SNAKE_TAIL_WINTER_2019_BONHOMME  = "bonhomme"
	SNAKE_TAIL_WINTER_2019_FLAKE     = "flake"
	SNAKE_TAIL_WINTER_2019_ICE_SKATE = "ice-skate"
	SNAKE_TAIL_WINTER_2019_PRESENT   = "present"

	// tails - stay home and code 2020
	SNAKE_TAIL_CODE_2020_COFEE      = "coffee"
	SNAKE_TAIL_CODE_2020_MOUSE      = "mouse"
	SNAKE_TAIL_CODE_2020_TIGER_TAIL = "tiger-tail"
	SNAKE_TAIL_CODE_2020_WEIGHT     = "weight"
)
