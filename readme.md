# CPU信息
> 提供获取CPU信息的接口

## vela.cpu
> 获取方法 返回summary类型数据
> 支持:windows 和 linux

### summary
> 汇总系统架构 CPU核心数 等信息 满足 index 接口


#### 字段
- total  &emsp;总时间周期
- num    &emsp;CPU核心数
- cnt    &emsp;物理核心数
- model  &emsp; CPU系列代号
- vendor &emsp;CPU制造商
- [sample](####采样信息)  采样信息

#### 函数
- [update()](####函数) &emsp;更新缓存信息
```lua
    vela.cpu.update()
    local cpu = vela.cpu
```

#### 采样信息
> stat = vela.cpu.sample 返回state <br />
> stat 字段除了name 其他都为浮点新也就是lua的number类型
- stat.name
- stat.total
- stat.user
- stat.system
- stat.idle
- stat.nice
- stat.io_wait
- stat.irq
- stat.soft_irq
- stat.steal
- stat.guest
- stat.guest_nice
```lua
    local cpu = vela.cpu
    cpu.update()

    local stat = cpu.sample
    print(stat.total)
    print(stat.user)
    print(stat.system)
    print(stat.idle)
    print(stat.nice)
    print(stat.io_wait)
    print(stat.irq)
    print(stat.soft_irq)
    print(stat.steal)
    print(stat.guest)
    print(stat.guest_nice)
```

#### example
```lua
    local cpu = vela.cpu
    cpu.update() --
    print(cpu.total)
    print(cpu.num)
    print(cpu.cnt)
    print(cpu.model)
    print(cpu.vendor)

```