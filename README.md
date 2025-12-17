---

##  Anti-Detection & Stealth Techniques

Implemented stealth mechanisms include:

1. Randomized timing & think delays  
2. Human-like mouse movement with overshoot & micro-corrections  
3. Human-like typing with variable speed and typo correction  
4. Natural scrolling behavior with pauses and reverse scrolling  
5. Activity scheduling (business hours only)  
6. Rate limiting & daily action quotas  
7. Browser fingerprint masking  
   - automation flags disabled  
   - randomized viewport dimensions  
   - `navigator.webdriver` overridden  
8. Duplicate action prevention via persisted state  

These techniques are combined to simulate **authentic human behavior patterns**.

---

## Authentication System (PoC)

- Credentials are loaded from environment variables
- Graceful handling of:
  - missing credentials
  - login failures
  - security checkpoints (2FA / captcha)
- **No security mechanisms are bypassed**
- Login flow exists purely as a **design demonstration**

---

## Search & Targeting (Mock Implementation)

- Query-based targeting:
  - job title
  - company
  - keywords
- Pagination handling
- Duplicate profile detection
- Safe mock profile URLs (no scraping)

This approach demonstrates **search architecture** without violating platform rules.

---

## Connection Requests (Mock)

- Simulated profile navigation
- Daily request limits enforced
- Human-like delays between actions
- Persistent tracking to prevent duplicate requests
- Clean separation of logic via connector module

---

## 💬 Messaging System (Mock)

- Template-based follow-up messages
- Dynamic variable substitution (e.g., name)
- Message history tracking
- Duplicate prevention across runs

---

## State Persistence

- JSON-based persistent storage
- Tracks:
  - sent connection requests
  - sent messages
- Allows safe resume after restarts
- Prevents repeated actions across executions

---

## Structured Logging

- Leveled logs:
  - INFO
  - WARN
  - ERROR
- Timestamped output
- Clear operational visibility during execution

---

## Setup Instructions

### 1. Install Go

https://go.dev/dl/

### 2. Clone Repository

```bash
git clone https://github.com/Vassu627/linkedin-automation
cd linkedin-automation-poc
```
