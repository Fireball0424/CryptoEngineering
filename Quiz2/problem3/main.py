from hashlib import sha256

# Logging
import logging
logger = logging.getLogger(__name__)
logging.addLevelName(logging.ERROR, 'EROR')
logging.basicConfig(filename='logger.log', encoding='utf-8', level=logging.DEBUG,
                    format="%(asctime)s [%(levelname)-4s] %(message)s",
                    datefmt="%Y/%m/%d %H:%M:%S", force=True)

# PreImage
studentID = "111550130"
preImage = sha256(studentID.encode())
logger.info(f"[preImage] {preImage.hexdigest()}")

# Miniming

prevBlockHash = preImage.hexdigest()

for Round in range(1, 10):
    # check whether need nonce
    NeedNonce = False
    for i in range(Round):
        if prevBlockHash[i] != studentID[i]:
            NeedNonce = True
            break

    if not NeedNonce:
        logger.info(f"[Round {Round} without nonce] {prevBlockHash}")
        Round += 1
        continue

    FindBlock = False
    for i in range(0x00000000, 0xffffffff+1):
        nonce = f"{i:08x}"
        block = prevBlockHash + nonce
        blockHash = sha256(block.encode())

        Find = True
        for b in range(Round):
            if blockHash.hexdigest()[b] != studentID[b]:
                Find = False
                break

        if Find:
            logger.info(f"[Round {Round} with nonce {nonce}] {blockHash.hexdigest()}")
            FindBlock = True
            prevBlockHash = blockHash.hexdigest()
            Round += 1
            break

    if not FindBlock:
        logger.error(f"[Round {Round}] not found with running out of nonce")
        break

