# 1BRC GO

## Data
The data was generated using the `createMeasurements` script from the 1BRC Repository.

## Read Results

The results of reading 16GB and a total of 1 billion lines of data and subsequent noop processing in a different goroutine.


### Scanner
- 74272098ms / 74.2720981s

### Scanner 1_000_000
- 74213911ms / 74.213911s

### Scanner 50_000_000
- 76272659ms / 76.2726595s

### Scanner 500_000_000
eats RAM like a chunky boy.
- 150064221ms / 150.064221s

Long time maybe due to paging.
Limiting factor is probably channel overhead since using `_ = scanner.Text()` reduces time with Buf 1000000 to 31s.


### Default Buf Reader
Without any optimizations
- 80516980 ms / 80.5169806 s

### Scanner 1 _000 Buffer 1_000 Lines Chunks
- 41649197 ms / 41.6491979 s

Using `batch = make([]string, DATA_CHAN_CHUNKSIZE)` instead of `batch = nil` tanks the performance of the program. Why exactly is unkown, i havent looked it up.

### File Chunking 1k Buffer 1MB IO Chunk 
- 9632183 ms / 9.6321835 s

But this has bugs because no new underlying arrays are beeing created.

- 13217743 ms / 13.2177437 s

Creating a new buffer on each iteration.

## Testing Env Specs
```
Processors
-------------------------------------------------------------------------

CPU Groups			1
CPU Group 0			20 threads, mask=0xFFFFF

Number of sockets		1
Number of threads		20

Processors Information
-------------------------------------------------------------------------

Socket 1			ID = 0
	Number of cores		14 (max 14)
	Number of threads	20 (max 20)
	Hybrid			yes, 2 coresets
	Core Set 0		P-Cores, 6 cores, 12 threads
	Core Set 1		E-Cores, 8 cores, 8 threads
	Manufacturer		GenuineIntel
	Name			Intel Core i7 13700H
	Codename		Raptor Lake
	Specification		13th Gen Intel(R) Core(TM) i7-13700H
	Package (platform ID)	Socket 1744 FCBGA (0x7)
	CPUID			6.A.2
	Extended CPUID		6.BA
	Core Stepping		J0
	Technology		10 nm
	TDP Limit		45.0 Watts
	Tjmax			100.0 C
	Core Speed		997.6 MHz
	Multiplier x Bus Speed	10.0 x 99.8 MHz
	Base frequency (cores)	99.8 MHz
	Base frequency (mem.)	99.8 MHz
	Stock frequency		0 MHz
	Max frequency		0 MHz
	Instructions sets	MMX, SSE, SSE2, SSE3, SSSE3, SSE4.1, SSE4.2, EM64T, AES, AVX, AVX2, FMA3, SHA
	Microcode Revision	0x4118
	L1 Data cache		6 x 48 KB (12-way, 64-byte line) + 8 x 32 KB (8-way, 64-byte line)
	L1 Instruction cache	6 x 32 KB (8-way, 64-byte line) + 8 x 64 KB (8-way, 64-byte line)
	L2 cache		6 x 1.25 MB (10-way, 64-byte line) + 2 x 2 MB (16-way, 64-byte line)
	L3 cache		24 MB (12-way, 64-byte line)
	Preferred cores		2 (#2, #3)
	Max CPUID level		00000020h
	Max CPUID ext. level	80000008h
	FID/VID Control		yes

Chipset
-------------------------------------------------------------------------

Northbridge			Intel Raptor Lake rev. 00
Southbridge			Intel Raptor Lake-P PCH rev. 01
Bus Specification		PCI-Express 4.0 (16.0 GT/s)
Graphic Interface		PCI-Express
Memory Type			DDR5
Memory Size			16 GBytes
Channels			4 x 32-bit
Memory Frequency		2593.7 MHz (1:26)
Memory Max Frequency		3733.3 MHz
CAS# latency (CL)		60.0
RAS# to CAS# delay (tRCD)	48
RAS# Precharge (tRP)		48
Cycle Time (tRAS)		112
Bank Cycle Time (tRC)		160
Row Refresh Cycle Time (tRFC)	728
Command Rate (CR)		1T
Uncore Frequency		2593.7 MHz
Memory Controller Frequency	1296.8 MHz

Storage
-------------------------------------------------------------------------

Drive	0/0/-1/-1
	Device Path		\\?\scsi#disk&ven_nvme&prod_mz9l4512hblu-00b#5&39f7ed2&0&000000#{53f56307-b6bf-11d0-94f2-00a0c91efb8b}
	Name			MZ9L4512HBLU-00BMV-SAMSUNG
	Revision		HXC75M0Q
	Capacity		476.9 GB
	Type			Fixed, SSD
	Bus Type		NVMe (17)
	Controller		NVM Express (NVMe) Controller at bus 2, device 0, function 0
	Link Speed		PCI-E 4x @ 16.0 GT/s

Software
-------------------------------------------------------------------------

Windows Version			Microsoft Windows 11  Home (x64), Version 23H2, Build 22631.3155
DirectX Version			12.0

```