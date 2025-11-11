
---

# 🧠 Heapless — Terminal Snake in Pure Go (No Heap, No AI)

A **terminal-based Snake game** written in **Go**, built under strict constraints:
**no heap**, **no AI**, **no dependencies** — only pure stack logic and terminal control.

This project explores **absolute control** over Go’s memory behavior and runtime.
It’s part game, part technical statement: *you don’t need a garbage collector to move a snake.*

---

## ⚙️ Rules of the Project

1. **🚫 No Heap Allocations**

   * Must be verified with:

     ```bash
     go build -gcflags="-m"
     ```

     If the output contains `"escapes to heap"`, it’s invalid.
     Everything must live on the stack.

     If fmt.Println or similar that gives out some heap warning assume it is for debugging only and it will be removed in up coming commits.


2. **🧍 No AI Assistance**

   * This README is the **only** exception — AI was used *only* to write this document.
   * From here on, **no ChatGPT**, **no Copilot**, **no LLMs**.
   * **Google** is allowed for documentation and research only.

3. **📦 No External Packages**

   * Only Go’s **standard library** is permitted.
   * No third-party modules, frameworks, or “helpers.”

4. **🧩 Stack Discipline**

   * Fixed-size arrays only.
   * No `append()`, no `make()` with dynamic length, no maps.
   * Every value must die where it was born.

---

## 🐍 Concept

Heapless is a minimal **Snake clone** that:

* Renders via **ANSI escape codes** (`\033` sequences).
* Handles real-time movement and collision manually.
* Operates entirely with **static memory**.
* Verifies memory safety using compiler analysis.

---

## 🧰 Run & Verify

Run the game:

```bash
go run .
```

Verify no heap allocations:

```bash
go build -gcflags="-m" .
```

If you see *“escapes to heap”*, you failed the rule. Fix it.

---

Perfect — I can expand that section with a deeper technical explanation, linking it to embedded systems and why avoiding the heap is an important learning exercise. Here’s a rewrite for your README section:

---

## 🧠 Why No Heap?

Heapless is a challenge in **predictable memory** and **deterministic behavior**.
By staying entirely on the stack:

* There’s **no GC interference** — the garbage collector doesn’t run, avoiding unpredictable pauses.
* **No runtime pauses** — your code executes in tight, deterministic loops.
* **No hidden allocations** — you know exactly where every byte lives and dies.
* You fully control every byte and lifetime, gaining **deep insight into memory management**.

Avoiding heap allocations is more than a technical stunt — it’s **essential for writing efficient, low-level, and embedded systems code**, where CPU cycles and memory usage are critical. On microcontrollers or performance-critical systems, dynamic allocation can:

* Cause fragmentation
* Introduce unpredictable latency
* Lead to runtime errors if memory runs out

By relying primarily on **stack-based memory**, you write code that is **faster, safer, and fully predictable**, which is exactly the mindset required in embedded systems development. Efficient embedded systems coding comes from **maximizing CPU usage while minimizing runtime overhead and heap dependency** ([source](https://www.embedded.com/why-embedded-software-must-avoid-dynamic-memory-allocation/)).

Learning to **avoid the heap in Go** is an excellent exercise because it forces you to:

* Think about memory lifetimes explicitly
* Use fixed-size arrays and static data structures
* Understand Go’s stack vs heap allocation model in depth

This not only improves **systems-level understanding** but also teaches **practical skills applicable to C, Rust, or any low-level programming environment**.

Learn more about Go’s memory model:
👉 [https://go.dev/blog/stack](https://go.dev/blog/stack)

---


## 🧩 Planned Features

* [ ] Terminal grid rendering with ANSI codes
* [ ] Non-blocking input
* [ ] Static board array
* [ ] Collision + growth logic
* [ ] Score display

---

## 🧾 License

MIT License — free to use, modify, and break as you please.

---
