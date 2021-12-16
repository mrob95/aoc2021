package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"math"
	"os"
)

const SUM = 0
const PRODUCT = 1
const MINIMUM = 2
const MAXIMUM = 3
const CONSTANT = 4
const GREATER = 5
const LESS = 6
const EQUAL = 7

type Packet struct {
	version, id int
	literal     int
	subs        []Packet
}

func parsePacket(in []bool, startPos int) (p Packet, endPos int) {
	pos := startPos
	result := Packet{}
	result.version = bitsToNumber(in[pos : pos+3])
	pos += 3
	result.id = bitsToNumber(in[pos : pos+3])
	pos += 3
	if result.id == CONSTANT {
		var values []int
		for in[pos] {
			values = append(values, bitsToNumber(in[pos+1:pos+5]))
			pos += 5
		}
		values = append(values, bitsToNumber(in[pos+1:pos+5]))
		pos += 5
		for i, byt := range values {
			result.literal |= (byt << ((len(values) - 1 - i) * 4))
		}
	} else {
		ltid := in[pos]
		pos += 1
		if ltid {
			numPackets := bitsToNumber(in[pos : pos+11])
			pos += 11
			for i := 0; i < numPackets; i++ {
				p, newPos := parsePacket(in, pos)
				result.subs = append(result.subs, p)
				pos = newPos
			}
		} else {
			subPacketLength := bitsToNumber((in[pos : pos+15]))
			pos += 15
			targetPos := pos + subPacketLength
			for pos != targetPos {
				p, newPos := parsePacket(in, pos)
				result.subs = append(result.subs, p)
				pos = newPos
			}
		}
	}
	return result, pos
}

func bytesToBits(in []byte) []bool {
	var result []bool
	for _, b := range in {
		for i := 7; i >= 0; i-- {
			bit := (b >> i) & 1
			result = append(result, bit == 1)
		}
	}
	return result
}

func bitsToNumber(in []bool) int {
	result := 0
	for i, b := range in {
		if b {
			result |= (0x1 << (len(in) - 1 - i))
		}
	}
	return result
}

func versionNumbers(p Packet) int {
	result := p.version
	for _, sub := range p.subs {
		result += versionNumbers(sub)
	}
	return result
}

func packetValue(p Packet) int {
	switch p.id {
	case SUM:
		result := 0
		for _, sub := range p.subs {
			result += packetValue(sub)
		}
		return result

	case PRODUCT:
		result := 1
		for _, sub := range p.subs {
			result *= packetValue(sub)
		}
		return result

	case MINIMUM:
		result := math.MaxInt64
		for _, sub := range p.subs {
			v := packetValue(sub)
			if v < result {
				result = v
			}
		}
		return result

	case MAXIMUM:
		result := math.MinInt64
		for _, sub := range p.subs {
			v := packetValue(sub)
			if v > result {
				result = v
			}
		}
		return result

	case CONSTANT:
		return p.literal

	case GREATER:
		first, second := packetValue(p.subs[0]), packetValue(p.subs[1])
		if first > second {
			return 1
		} else {
			return 0
		}

	case LESS:
		first, second := packetValue(p.subs[0]), packetValue(p.subs[1])
		if first < second {
			return 1
		} else {
			return 0
		}

	case EQUAL:
		first, second := packetValue(p.subs[0]), packetValue(p.subs[1])
		if first == second {
			return 1
		} else {
			return 0
		}
	}
	panic("Unreachable")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	bytes, _ := hex.DecodeString(line)
	bits := bytesToBits(bytes)

	p, _ := parsePacket(bits, 0)

	fmt.Println(versionNumbers(p))
	fmt.Println(packetValue(p))
}
