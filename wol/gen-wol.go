import (
    "fmt"
    "net"
    "strconv"
)

func send(macAddress string) (error) {
    hwBytes, err := net.ParseMAC(macAddress)
    if err != nil {
        return nil, err
    }

	macBytes [6]byte

	for idx := range mac {
		macBytes[idx] = hwBytes[idx]
	} 

    // Create the magic packet as a byte array
    packet := []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff}
    for i := 0; i < 16; i++ {
        packet = append(packet, macBytes...)
    }

    // Create a UDP broadcast address
    udpAddr, err := net.ResolveUDPAddr("udp", "255.255.255.255:9")
    if err != nil {
        return nil, err
    }

    // Create a UDP connection
    conn, err := net.DialUDP("udp", nil, udpAddr)
    if err != nil {
        return nil, err
    }
    defer conn.Close()

    // Transmit the magic packet
    _, err = conn.Write(packet)
    if err != nil {
        return nil, err
    }
}