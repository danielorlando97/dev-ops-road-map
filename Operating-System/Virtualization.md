## Virtualization

Virtualization is the creation of a virtual -- rather than actual -- version
of something, such as an operating system (OS), a server, a storage device
or network resources. It uses software that simulates hardware functionality
to create a virtual system. This practice allows IT organizations to operate
multiple operating systems, more than one virtual system and various
applications on a single server.

- What is Virtualization? and its types?
- What is Hypervisor and VM?
- Containers vs VM

```bash
    go build -o go.program ./examples/virtual-process.example.go
    ./go.program A & ; ./go.program B & ; ./go.program C & ; ./go.program D & ;
```

Well, now things are getting a little more interesting. Even though we
have only one processor, somehow all four of these programs seem to be
running at the same time! How does this magic happen?4
It turns out that the operating system, with some help from the hardware
, is in charge of this illusion, i.e., the illusion that the system has a
very large number of virtual CPUs. Turning a single CPU (or small set of
them) into a seemingly infinite number of CPUs and thus allowing many
programs to seemingly run at once is what we call virtualizing the CPU,
the focus of the first major part of this book.

```bash
    go build -o go.program ./examples/virtual-memory.example.go
    ./go.program & ; ./go.program &
```

### Basic Notes

A key use of virtualization technology is server virtualization, which uses
a software layer -- called a hypervisor -- to emulate the underlying hardware.
This often includes the CPU's memory, input/output (I/O) and network traffic.

What is a hypervisor?

A hypervisor is a program for creating and running virtual machines.
Hypervisors have traditionally been split into two classes: type one, or
"bare metal" hypervisors that run guest virtual machines directly on a
system's hardware, essentially behaving as an operating system. Type two,
or "hosted" hypervisors behave more like traditional applications that can
be started and stopped like a normal program. In modern systems, this split
is less prevalent, particularly with systems like KVM. KVM, short for
kernel-based virtual machine, is a part of the Linux kernel that can run
virtual machines directly, although you can still use a system running KVM
virtual machines as a normal computer itself

What is a virtual machine?

A virtual machine is the emulated equivalent of a computer system that runs
on top of another system. Virtual machines may have access to any number of
resources: computing power, through hardware-assisted but limited access to
the host machine's CPU and memory; one or more physical or virtual disk devices
for storage; a virtual or real network inferface; as well as any devices such
as video cards, USB devices, or other hardware that are shared with the virtual
machine. If the virtual machine is stored on a virtual disk, this is often
referred to as a disk image. A disk image may contain the files for a virtual
machine to boot, or, it can contain any other specific storage needs.

The virtualization process follows the steps listed below:

- Hypervisors detach the physical resources from their physical environments.
- Resources are taken and divided, as needed, from the physical environment
  to the various virtual environments.
- System users work with and perform computations within the virtual environment.
- Once the virtual environment is running, a user or program can send an
  instruction that requires extra resources form the physical environment. In
  response, the hypervisor relays the message to the physical system and stores
  the changes. This process will happen at an almost native speed.
- The virtual environment is often referred to as a guest machine or virtual
  machine. The VM acts like a single data file that can be transferred from one
  computer to another and opened in both; it is expected to perform the same
  way on every computer.

#### Types of virtualization

- [Network virtualization](https://searchservervirtualization.techtarget.com/definition/network-virtualization)
- [Storage virtualization](https://www.techtarget.com/searchstorage/definition/storage-virtualization)
- [Data virtualization](https://www.techtarget.com/searchdatamanagement/definition/data-virtualization)
- [Desktop virtualization](https://www.techtarget.com/searchvirtualdesktop/definition/desktop-virtualization)
- [Application virtualization](https://www.techtarget.com/searchvirtualdesktop/definition/app-virtualization)
- [Server virtualization](https://searchservervirtualization.techtarget.com/definition/server-virtualization):
  Is the masking of server resources -- including the number and identity of
  individual physical servers, processors and operating systems -- from server users.
  The intention is to spare the user from having to understand and manage complicated
  details of server resources while increasing resource sharing and utilization and
  maintaining the capacity to expand later.
  The layer of software that enables this abstraction is often referred to as the
  hypervisor. The most common hypervisor -- Type 1 -- is designed to sit directly
  on bare metal and provide the ability to virtualize the hardware platform for
  use by the virtual machines. KVM virtualization is a Linux kernel-based
  virtualization hypervisor that provides Type 1 virtualization benefits like other
  hypervisors. KVM is licensed under open source. A Type 2 hypervisor requires a
  host operating system and is more often used for testing and labs.

##### Advantages and Benefits [here](https://www.techtarget.com/searchitoperations/definition/virtualization)

- Lower costs
- Easier disaster recovery
- Easier testing
- Quicker backups
- Improved productivity

---

- Single-minded servers
- Expedited deployment and redeployment
- Reduced heat and improved energy savings
- Better for the environment
- Lack of vendor dependency
- Easier migration to the cloud

##### [Read About Limitations](https://www.techtarget.com/searchitoperations/definition/virtualization)

##### [Server virtualization in DevOps continues to offer advantages](https://www.techtarget.com/searchitoperations/tip/Server-virtualization-in-DevOps-continues-to-offer-advantages)

### Containers vs. virtual machines

Containers and virtual machines are very similar resource virtualization technologies.
Virtualization is the process in which a system singular resource like RAM, CPU, Disk,
or Networking can be ‘virtualized’ and represented as multiple resources. The key
differentiator between containers and virtual machines is that virtual machines virtualize
an entire machine down to the hardware layers and containers only virtualize software
layers above the operating system level.

What is a container?

Containers are lightweight software packages that contain all the dependencies required
to execute the contained software application. These dependencies include things like
system libraries, external third-party code packages, and other operating system level
applications. The dependencies included in a container exist in stack levels that
are higher than the operating system.

### Advances Information

The OS creates this illusion by virtualizing the CPU. By running one
process, then stopping it and running another, and so forth, the OS can
promote the illusion that many virtual CPUs exist when in fact there is
only one physical CPU (or a few). This basic technique, known as time
sharing of the CPU, allows users to run as many concurrent processes as
they would like; the potential cost is performance, as each will run more
slowly if the CPU(s) must be shared.
To implement virtualization of the CPU, and to implement it well, the
OS will need both some low-level machinery as well as some high-level
intelligence

The abstraction provided by the OS of a running program is something
we will call a process.Instructions lie in memory; the data that the
running program reads and writes sits in memory as well. Thus the memory that the
process can address (called its address space) is part of the process
Note that there are some particularly special registers that form part
of this machine state. For example, the program counter (PC) (sometimes
called the instruction pointer or IP) tells us which instruction of the
program is currently being executed; similarly a stack pointer and associated
frame pointer are used to manage the stack for function parameters, local
variables, and return addresses.
Finally, programs often access persistent storage devices too. Such I/O
information might include a list of the files the process currently has open.

Operating systems are replete with various important data structures
that we will discuss in these notes. The process list is the first such structure.
It is one of the simpler ones, but certainly any OS that has the ability
to run multiple programs at once will have something akin to this structure in
order to keep track of all the running programs in the system.
Sometimes people refer to the individual structure that stores information
about a process as a Process Control Block (PCB), a fancy way of
talking about a C structure that contains information about each process.
