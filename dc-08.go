package main

import (
	"encoding/asn1"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	target := "192.168.11.1"
	community := "public"

	// SNMP GET request PDU
	pdu := buildSnmpGetRequest()

	// Build SNMP message
	snmpMessage := buildSnmpMessage(community, pdu)

	// Send SNMP message and receive response
	response, err := sendSnmpRequest(target, snmpMessage)
	if err != nil {
		log.Fatal("Error sending SNMP request:", err)
		os.Exit(1)
	}

	// Decode and print SNMP response
	result, err := decodeSnmpResponse(response)
	if err != nil {
		log.Fatal("Error decoding SNMP response:", err)
		os.Exit(1)
	}

	fmt.Printf("Response: %v\n", result)
}

// buildSnmpGetRequest builds an SNMP GET request PDU
func buildSnmpGetRequest() []byte {
	// SNMP GET request PDU
	pdu := []byte{0x30, 0x06, 0x02, 0x01, 0x01, 0x05, 0x00} // Example OID for system description

	return pdu
}

// buildSnmpMessage builds an SNMP message with the specified community and PDU
func buildSnmpMessage(community string, pdu []byte) []byte {
	// SNMP message
	message := []byte{
		0x30, 0x2F, // Sequence (Message)
		0x02, 0x01, 0x01, // SNMP version 2c
		0x04, byte(len(community)), // Community string length
	}

	message = append(message, []byte(community)...) // Community string
	message = append(message,
		0xA0, byte(len(pdu)), // SNMP PDU type (GET)
	)

	message = append(message, pdu...) // PDU

	return message
}

// sendSnmpRequest sends an SNMP request to the target and returns the response
func sendSnmpRequest(target string, message []byte) ([]byte, error) {
	conn, err := net.DialTimeout("udp", target+":161", 2*time.Second)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	_, err = conn.Write(message)
	if err != nil {
		return nil, err
	}

	response := make([]byte, 4096)
	_, err = conn.Read(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// decodeSnmpResponse decodes an SNMP response and returns the result
func decodeSnmpResponse(response []byte) (interface{}, error) {
	var result interface{}
	_, err := asn1.Unmarshal(response, &result)
	return result, err
}
