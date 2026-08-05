# LeetCode Design Solutions

[![Stars](https://img.shields.io/github/stars/cloudnative0x0/leetcode-design?style=flat-square)](https://github.com/cloudnative0x0/leetcode-design/stargazers)
[![Forks](https://img.shields.io/github/forks/cloudnative0x0/leetcode-design?style=flat-square)](https://github.com/cloudnative0x0/leetcode-design/network/members)
[![Go Version](https://img.shields.io/badge/Go-1.24.2-blue?style=flat-square&logo=go)](https://go.dev/)
[![License](https://img.shields.io/github/license/cloudnative0x0/leetcode-design?style=flat-square)](LICENSE)
[![Last Commit](https://img.shields.io/github/last-commit/cloudnative0x0/leetcode-design?style=flat-square)](https://github.com/cloudnative0x0/leetcode-design/commits/main)

[English](#english) | [Русский](#русский)

> “A ship in port is safe, but that's not what ships are built for.” – Grace Hopper

---

## English

### About
Collection of Go implementations for LeetCode Design problems.  
Every solution lives in its own folder and ships with stress tests (randomised, 200k+ operations).

### Problems

| # | Title | Difficulty | Time | Space |
|---|-------|------------|------|-------|
| 146 | LRU Cache | Hard | O(1) | O(capacity) |
| 155 | Min Stack | Easy | O(1) | O(n) |
| 225 | Stack using Queues | Easy | Push O(n) | O(n) |
| 232 | Queue using Stacks | Easy | amortized O(1) | O(n) |
| 460 | LFU Cache | Hard | O(1) | O(capacity) |
| 706 | Design HashMap | Easy | avg O(1) | O(n) |

### How to run
```bash
go test -race ./...                        # entire repo
cd 232_queue_using_stacks && go test -v    # single problem
```

### Authors
by Herman Murauyou: [LeetCode](https://leetcode.com/CloudNative0x0) · [Codeforces](https://codeforces.com/profile/cloudstrategist) · [LinkedIn](https://www.linkedin.com/in/cloudnative0x0)

---

## Русский

> «Корабль в гавани в безопасности, но не для этого строят корабли.» – Грейс Хоппер

### О проекте
Решения задач раздела Design с LeetCode на Go.  
В каждой папке – исходный код и стресс-тесты (рандомизированные, от 200 тысяч операций).

### Задачи

| # | Название | Сложность | Время | Память |
|---|----------|-----------|-------|--------|
| 146 | LRU Cache | Hard | O(1) | O(capacity) |
| 155 | Min Stack | Easy | O(1) | O(n) |
| 225 | Stack using Queues | Easy | Push O(n) | O(n) |
| 232 | Queue using Stacks | Easy | амортизированное O(1) | O(n) |
| 460 | LFU Cache | Hard | O(1) | O(capacity) |
| 706 | Design HashMap | Easy | среднее O(1) | O(n) |

### Запуск
```bash
go test -race ./...                        # всё сразу
cd 232_queue_using_stacks && go test -v    # одна задача
```

### Авторы
by Herman Murauyou: [LeetCode](https://leetcode.com/CloudNative0x0) · [Codeforces](https://codeforces.com/profile/cloudstrategist) · [LinkedIn](https://www.linkedin.com/in/cloudnative0x0)