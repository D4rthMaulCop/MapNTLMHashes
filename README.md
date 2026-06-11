# MapNTLMHashes

A tool for matching cracked NTLM password hashes against an NTDS dump file, mapping plaintext passwords back to Active Directory accounts.

## Overview

During a penetration test, after obtaining an NTDS.dit dump and cracking hashes with a tool like Hashcat, the cracked output (potfile) contains `hash:password` pairs. This tool correlates those cracked hashes back to the original NTDS dump entries, producing a combined output with usernames, RIDs, hashes, and plaintext passwords.

## Usage

```
./MapNTLMHashes <potfile> <ntds-dump>
```

**Arguments:**
- `potfile` — Hashcat `.pot` file or similarly formatted file with `hash:password` entries
- `ntds-dump` — NTDS dump in secretsdump format (`username:RID:LM:NTLM:::`)

**Example:**
```
./MapNTLMHashes hashcat.potfile ntds.dit.ntds
```

**Output:**

Each matched line from the NTDS dump is printed with the cracked password appended:
```
Administrator:500:aad3b435b51404eeaad3b435b51404ee:31d6cfe0d16ae931b73c59d7e0c089c0::::<password>
[+] Total Matches: 42
```

## Building

#### AMD64
```bash
# compile for windows amd64
GOOS=windows GOARCH=amd64 go build -o myapp.exe

# compile for windows amd64 with CGO (make sure mingw-w64 is installed)
GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC="x86_64-w64-mingw32-gcc" CXX="x86_64-w64-mingw32-g++" go build -o myapp.exe
```
#### ARM64
```bash
# compile for linux amd64
GOOS=linux GOARCH=amd64 go build -o myapp

# compile for linux arm64
GOOS=linux GOARCH=arm64 go build -o myapp

# compile for windows arm64
GOOS=windows GOARCH=arm64 go build -o myapp.exe
```


Requires Go 1.x or later.

## Input Formats

**Hashcat Potfile** (`hash:password` per line):
```
31d6cfe0d16ae931b73c59d7e0c089c0:Password123
...
```

**NTDS dump** (secretsdump format):
```
DOMAIN\username:RID:LM_hash:NTLM_hash:::
```

## Use Case

Intended for authorized penetration testing and Active Directory security assessments. Helps quickly identify which accounts have been compromised by mapping cracked hashes to specific users for remediation reporting.
