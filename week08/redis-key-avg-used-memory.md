| Value Size (byte) | Used Memory | Increment |
| ----------------- | ----------- | --------- |
| 10                | 8227200     | 0%        |
| 20                | 9048080     | 9.98%     |
| 50                | 12248080    | 35.37%    |
| 100               | 17848800    | 45.73%    |
| 200               | 29048800    | 62.75%    |
| 1000              | 109048416   | 375.40%   |
| 5000              | 518648888   | 475.61%   |




```bash
debug populate 100000 10b 10
# before
used_memory:874144
used_memory_human:853.66K
# after
used_memory:9101344
used_memory_human:8.68M
```

```bash
debug populate 100000 20b 20
# before
used_memory:853728
used_memory_human:833.72K
# after
used_memory:9901808
used_memory_human:9.44M
```

```bash
debug populate 100000 50b 50
# before
used_memory:854192
used_memory_human:834.17K
# after
used_memory:13102272
used_memory_human:12.50M
```

```bash
debug populate 100000 100b 100
# before
used_memory:855120
used_memory_human:835.08K
# after
used_memory:18703920
used_memory_human:17.84M
```

```bash
debug populate 100000 200b 200
# before
used_memory:854656
used_memory_human:834.62K
# after
used_memory:29903456
used_memory_human:28.52M
```

```bash
debug populate 100000 1000b 1000
# before
used_memory:856056
used_memory_human:835.99K
# after
sed_memory:109904472
used_memory_human:104.81M

```

```bash
debug populate 100000 5000b 5000
# before
used_memory:856056
used_memory_human:835.99K
# after
used_memory:519504944
used_memory_human:495.44M
```



