# Critique2 

## How to run 

```bash
go run main.go
```

You can see the interface in http://localhost:8080

## File Structure 
```
- Critique2/
|- block/
    |- block.go         // Basic Operation of blocks
    |- Multiply.go      // Algorithm 1 
    |- block_test.go    // Test for basic operation of blocks and multiply
|- gcm/ 
    |- GHASH.go         // Algorithm 2
    |- GCTR.go          // Algorithm 3
    |- gcm.go           // Algorithm 4 and 5 
    |- gcm_test.go      // Test GCTR, encryption and decryption functions 
|- utils/ 
    |- utils.go         // Bits Basic Operations and Inc32 
    |- ciph.go          // CIPH_K (know use AES) 
    |- const.go         // Specific Key Length and TagLength 
    |- utils_test.go    // Test convertion between bits and string 
|- templates/
    |- index.html       // Front-end page 
|- main.go              // Can modify here to test different data 
```

## Web Input Format
* PlainText: String 
* AAD: String 
* IV: String 

* CipherText: HexString
* Tags: HexString 