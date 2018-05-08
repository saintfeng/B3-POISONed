package main

import(
    "net"
    "log"
    "fmt"
    "strconv"
    // "math/big"
    "encoding/json"
    "time"
    
    "github.com/bytom/testutil"
    "github.com/bytom/protocol/bc/types"
    // "github.com/bytom/consensus/difficulty"
    // "github.com/bytom/consensus"
)

type t_err struct {
    Code            uint64      `json:"code"`
    Message         string      `json:"message"`
}

type t_job struct {
    Version         string      `json:"version"`
    Height          string      `json:"height"`
    PreBlckHsh      string      `json:"previous_block_hash"`
    Timestamp       string      `json:"timestamp"`
    TxMkRt          string      `json:"transactions_merkle_root"`
    TxStRt          string      `json:"transaction_status_hash"`
    Nonce           string      `json:"nonce"`
    Bits            string      `json:"bits"`
    JobId           string      `json:"job_id"`
    Seed            string      `json:"seed"`
    Target          string      `json:"target"`
}

type t_result struct {
    Id              string      `json:"id"`
    Job             t_job       `json:"job"`
    Status          string      `json:"status"`
}

type t_resp struct {
    Id              int64       `json:"id"`
    Jsonrpc         string      `json:"jsonrpc, omitempty"`
    Result          t_result    `json:"result, omitempty"`
    Error           t_err       `json:"error, omitempty"`
}

const (
    maxNonce = ^uint64(0) // 2^64 - 1 = 18446744073709551615
    poolAddr = "stratum-btm.antpool.com:6666" //39.107.125.245
    // poolAddr = "192.168.199.189:5555"

    flush = "\r\n\r\n"
    DEBUG = false
    esHR     = 80 //estimated Hashrate
)

var (
    lastNonce  = ^uint64(0)
    // lastHeight = "-1"
    Id = uint64(0)
)

func main() {
start:
    Id += 1
    conn, err := net.Dial("tcp", poolAddr)
    if err != nil {
        log.Fatalln(err)
    }
    defer conn.Close()

    send_msg := `{"method": "login", "params": {"login": "haoyuyu.1", "pass": "123", "agent": "bmminer/2.0.0"}, "id": `
    send_msg += strconv.FormatUint(Id, 10)
    send_msg += `}`
    // send_msg := `{"method": "login", "params": {"login": "bm1qcxg0w7c70tdd46t7dxn204mkyeyudcz063s49e", "pass": "123", "agent": "bmminer/2.0.0"}, "id": 1}`
    conn.Write([]byte(send_msg))
    conn.Write([]byte(flush))
    log.Printf("Sent: %s", send_msg)

    buff := make([]byte, 1024)
    n, _ := conn.Read(buff)
    log.Printf("Received: %s", buff[:n])

    var resp t_resp
    json.Unmarshal([]byte(buff[:n]), &resp)
    if DEBUG {
        mock_input(&resp)
    }

    // if lastHeight != resp.Result.Job.Height {
    //     lastNonce = str2ui64Bg(resp.Result.Job.Nonce)-1
    // }
    // lastNonce = str2ui64Bg(resp.Result.Job.Nonce)-1
    lastNonce = str2ui64Li(resp.Result.Job.Nonce)-1
    // lastHeight = resp.Result.Job.Height
    // if !mine(resp.Result.Job) {
    //     goto start
    // }
    mine(resp.Result.Job)
    nonce := lastNonce
    nonceStr := strconv.FormatUint(nonce, 16)
    // nonceStr = strSwitchEndian(fmt.Sprintf("%016s", nonceStr))
    nonceStr = fmt.Sprintf("%016s", nonceStr)
    if DEBUG {
        log.Printf("sending back nonce as string: %s", nonceStr)
    }

    send_msg = `{"method": "submit", "params": {"id": "haoyuyu.1", "job_id": "`
    send_msg += resp.Result.Job.JobId
    send_msg += `", "nonce": "`
    send_msg += nonceStr
    send_msg += `"}, "id":`
    send_msg += strconv.FormatUint(Id, 10)
    send_msg += `}`
    conn.Write([]byte(send_msg))
    conn.Write([]byte(flush))
    log.Printf("Sent: %s", send_msg)
    buff = make([]byte, 1024)
    n, _ = conn.Read(buff)
    log.Printf("Received: %s\n", buff[:n])
    json.Unmarshal([]byte(buff[:n]), &resp)
    goto start
}

/*
type BlockHeader struct {
    Version             uint64  // The version of the block.
    Height              uint64  // The height of the block.
    PreviousBlockHash   bc.Hash // The hash of the previous block.
    Timestamp           uint64  // The time of the block in seconds.
    Nonce               uint64  // Nonce used to generate the block.
    Bits                uint64  // Difficulty target for the block.
    BlockCommitment     types.BlockCommitment{
                            TransactionsMerkleRoot: node.transactionsMerkleRoot,
                            TransactionStatusHash:  node.transactionStatusHash,
                        },
}
*/

func mine(job t_job) bool {
    bh := &types.BlockHeader{
                Version:            str2ui64Bg(job.Version),
                Height:             str2ui64Bg(job.Height),
                PreviousBlockHash:  testutil.MustDecodeHash(job.PreBlckHsh),
                Timestamp:          str2ui64Bg(job.Timestamp),
                Bits:               str2ui64Bg(job.Bits),
                BlockCommitment:    types.BlockCommitment{
                                        TransactionsMerkleRoot: testutil.MustDecodeHash(job.TxMkRt),
                                        TransactionStatusHash:  testutil.MustDecodeHash(job.TxStRt),
                                    },
        }
    if DEBUG {
        view_parsing(bh, job)
    }

    lastNonce += 10000000
    time.Sleep(3 * time.Second)
    // log.Printf("Mining at height:\t%d\n", bh.Height)
    // log.Printf("Start from nonce:\t0x%016x = %d\n", lastNonce+1, lastNonce+1)
    // for i := str2ui64Bg(job.Nonce); i <= maxNonce; i++ {
    // for i := uint64(lastNonce + 1); i <= uint64(lastNonce+consensus.TargetSecondsPerBlock*esHR) && i <= maxNonce; i++ {
        // log.Printf("Checking PoW with nonce: 0x%016x = %d\n", i, i)
        // bh.Nonce = i
        // headerHash := bh.Hash()
        // if DEBUG {
        //     fmt.Println("headerHash:", headerHash.String())
        // }

        if true {
        // seedHash := testutil.MustDecodeHash(job.Seed)
        // if difficulty.CheckProofOfWork(&headerHash, &seedHash, bh.Bits) {
            // log.Printf("Block mined! Proof hash: 0x%v\n", headerHash.String())
            return true
        }
    // }
    // log.Printf("Stop at nonce:\t0x%016x = %d\n", bh.Nonce, bh.Nonce)
    lastNonce = bh.Nonce

    return false
}

func mock_input(presp *t_resp) {
    body := `{
                "id":1,
                "jsonrpc":"2.0",
                "result":{
                    "id":"antminer_1",
                    "job":{
                        "version":"0100000000000000",
                        "height":"0000000000000000",
                        "previous_block_hash":"0000000000000000000000000000000000000000000000000000000000000000",
                        "timestamp":"e55a685a00000000",
                        "transactions_merkle_root":"237bf77df5c318dfa1d780043b507e00046fec7f8fdad80fc39fd8722852b27a",
                        "transaction_status_hash":"53c0ab896cb7a3778cc1d35a271264d991792b7c44f5c334116bb0786dbc5635",
                        "nonce":"1055400000000000",
                        "bits":"ffff7f0000000020",
                        "job_id":"16942",
                        "seed":"8636e94c0f1143df98f80c53afbadad4fc3946e1cc597041d7d3f96aebacda07",
                        "target":"c5a70000"
                    },
                    "status":"OK"
                },
                "error":null
            }`
    json.Unmarshal([]byte(body), &(*presp))
}

func view_parsing(bh *types.BlockHeader, job t_job) {
    log.Println("Printing parsing result:")
    fmt.Println("\tVersion:", bh.Version)
    fmt.Println("\tHeight:", bh.Height)
    fmt.Println("\tPreviousBlockHash:", bh.PreviousBlockHash.String())
    fmt.Println("\tTimestamp:", bh.Timestamp)
    fmt.Println("\tbits_str:", job.Bits)
    fmt.Println("\tBits:", bh.Bits)
    fmt.Println("\tTransactionsMerkleRoot:", bh.BlockCommitment.TransactionsMerkleRoot.String())
    fmt.Println("\tTransactionStatusHash:", bh.BlockCommitment.TransactionStatusHash.String())
    fmt.Println("\ttarget_str:", job.Target)
    fmt.Println("\ttarget_ui64Bg:", str2ui64Bg(job.Target))
}

func str2ui64Bg(str string) uint64 {
    ui64, _ := strconv.ParseUint(strSwitchEndian(str), 16, 64)
    return ui64
}

func str2ui64Li(str string) uint64 {
    ui64, _ := strconv.ParseUint(str, 16, 64)
    return ui64
}

func strSwitchEndian(oldstr string) string {
    // fmt.Println("old str:", oldstr)
    slen := len(oldstr)
    if slen%2 != 0 {
        panic("hex string format error")
    }

    newstr := ""
    for i := 0; i < slen; i+=2 {
        newstr += oldstr[slen-i-2:slen-i]
    }
    // fmt.Println("new str:", newstr)
    return newstr
}
