package main

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"

	"fmt"
	_ "net/http/pprof"
)

func test(prefix string) {
	time.Sleep(10 * time.Second)

	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		fmt.Println(err)
		return
	}

	i := 0
	for {
		conn.Do("SET", "CONFIG:"+prefix+":"+strconv.Itoa(i), "jfdsjfkljfldjflfdafdafdafdsafdsafdsaf"+
			"kdsajfoudafodsafdsfkldsjafkdsnmvnafkljdsfljdsalkfjdlskajfldjfakldjsalfkjdsalkfjdsklafjldaffdkdfdsafdsafsdafafadsfasj"+
			"fldsajflkdsjaflkdjsalfkjdsalfjdlsajfldskjafkldsjalfkjdakfjdklsakfjdlksajflkdsajfdksafdsafdsafdsafdafdsafajfksadl"+
			"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
			"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
			"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
			"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
			"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
			"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
			"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
			"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
			"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
			"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
			"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
			/*			"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+
						"jflkdsjaflkdsajflkjdsakfjdslkafjdslkajfkdsajflafjdlskajflkdsajfklsadjfkldsjfsajfjfdklsafdsafdsafdsafdsafdsafdsafsafjsdklafjdska"+ */
			"fjdjfklsajfldafsafjdsalfjdskafjdslffdafdsafdafdsafdasfdsafdsafdaf", "EX", 2)
		//fmt.Println("set " + "CONFIG:" + strconv.Itoa(i))
		time.Sleep(5 * time.Millisecond)
		i++
	}
}

func main() {
	//n := runtime.GOMAXPROCS(12)
	//fmt.Println("n = ", n)
	go http.ListenAndServe("0.0.0.0:6060", nil) // add pprof

	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = conn.Do("CONFIG", "SET", "notify-keyspace-events", "KEA")
	if err != nil {
		fmt.Println(err)
		return
	}

	pid := os.Getpid()

	go test("00")

	go test("01")
	/*	go test("02")

		go test("03")
		go test("04")
		go test("05")
		go test("06")
		go test("07")
		go test("08")
		go test("09")
		go test("10")

			go test("11")
			go test("12")
			go test("13")
			go test("14")
			go test("15")
			go test("16")
			go test("17")
			go test("18")
			go test("19")
			go test("20")
			go test("21")
			go test("22")
			go test("23")
			go test("24")
			go test("25")
			go test("26")
			go test("27")
			go test("28")
			go test("29")
			go test("30")
			go test("31")*/

	psc := redis.PubSubConn{Conn: conn}
	psc.PSubscribe("__key*__:CONFIG:*")
	for {
		switch msg := psc.Receive().(type) {
		case redis.Message:
			fmt.Printf("pid: %d, Message: %s %s\n", pid, msg.Channel, msg.Data)
		//case redis.PMessage:
		//	fmt.Printf("PMessage: %s %s %s\n", msg.Pattern, msg.Channel, msg.Data)
		case redis.Subscription:
			fmt.Printf("pid: %d, Subscription: %s %s %d\n", pid, msg.Kind, msg.Channel, msg.Count)
			if msg.Count == 0 {
				return
			}
		case error:
			fmt.Printf("error: %v\n", msg)
			return
		}
	}
}
