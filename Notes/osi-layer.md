# OSI Model: An Overview

**Rachelle Miller**  
*Copyright SANS Institute 2021. Author Retains Full Rights.*  
*This paper was published by SANS Institute. Reposting is not permitted without express written permission.*

**Key fingerprint:** AF19 FA27 2F94 998D FDB5 DE3D F8B5 06E4 A169 4E46

## Introduction

The Open Systems Interconnection (OSI) reference model has been a foundational framework for computer networking since its inception in 1984. Proposed by the International Standards Organization (ISO), the OSI model aimed to establish design standards for equipment manufacturers to enable effective communication. It defines a hierarchical architecture that logically divides the functions necessary for system-to-system communication.

## OSI Model Overview

The OSI model consists of seven layers, each serving distinct functions with varying levels of abstraction. The principles guiding the model's creation include the need for different levels of abstraction, well-defined functions, standardized protocols, minimized information flow across interfaces, and a balance between the number of layers and architectural manageability.

### The Seven Layers

1. **Physical Layer**: This layer deals with raw bit transmission over communication channels. It defines transmission media, network topology, and the physical hardware's role.

2. **Data Link Layer**: Responsible for error control and synchronization in data transmission. It frames data, assigns source and destination addresses, and manages data flow control.

3. **Network Layer**: Handles routing, packet forwarding, and subnet control. It establishes end-to-end connections and controls congestion within the network.

4. **Transport Layer**: Manages end-to-end communication control, error checking, and flow control. It ensures accurate and reliable data transmission.

5. **Session Layer**: Facilitates ongoing communication sessions between applications across a network. It handles session setup, data exchanges, and tear-down procedures.

6. **Presentation Layer**: Manages data format conversion, syntax, and semantics during network communication. It deals with encryption, text compression, and reformatting.

7. **Application Layer**: Provides interfaces for applications to access network services. It offers services such as file transfer, virtual terminals, electronic mail, and directory services.

### OSI Layer Characteristics

The layered approach offers advantages like modular problem-solving and extensibility. The upper layers (application, presentation, and session) focus on application-specific functions, while the lower layers (transport, network, data link, and physical) provide primitive network-specific functions.

## Detailed Layer Descriptions

### Application Layer (Layer 7)

The application layer serves as the interface for users and computers to access networked services. It handles services like file transfer, virtual terminals, electronic mail, messaging handling, and directory services. Distributed application services share characteristics like end-user interfaces, mass storage input/output, and data transfer mechanisms.

### Presentation Layer (Layer 6)

The presentation layer manages data formatting and conversion for network communications. It handles syntax and semantics of transmitted information. Data compression, cryptography, and common communication services are provided here.

### Session Layer (Layer 5)

The session layer enables ongoing communications between parties in a session across a network. It manages session setup, data exchanges, and tear-down procedures. Security services and session identification control access to sessions.

### Transport Layer (Layer 4)

The transport layer ensures end-to-end communication control, error checking, and reliability. It corrects errors during transmission and manages acknowledgments. It can establish multiple network connections for improved throughput.

### Network Layer (Layer 3)

The network layer controls sub-net operation, routing, congestion control, and accounting. It handles both connectionless and connection-oriented services. The IP protocol resides in this layer, and routers operate at this level.

### Data Link Layer (Layer 2)

The data link layer transforms raw transmissions into error-free network-layer data. It breaks data into frames, assigns source/destination addresses, and provides error control and synchronization.

### Physical Layer (Layer 1)

The physical layer transmits raw bits over communication channels. It deals with transmission media, network topology, and connectivity.

## Conclusion

The OSI model, with its seven-layer structure, provides a conceptual framework for understanding computer networking. While not all networks fully utilize every layer, the model serves as a guide for protocol designers. It fosters a modular understanding of network components and their interactions, facilitating effective communication among networked computers.

**References:**

1. Feig, Rami. “Computer Networks: The OSI Reference Model.”
2. Mitchell, Bradley. “Basic Network Design – The OSI Model.”
3. Tan Teng Hong, Andrew; Chee Meng, Mah; Yew Wai, Chee; Yoke Chuan, Tan; Kim Ming, Cheong. “Comparing OSI and TCP/IP.”
4. Anderson, Christa. “Mapping Practice to Theory: NT Networking and the OSI Model.”
5. Briscoe, Neil. “Understanding The OSI 7-Layer Model.”
6. Zhou, Tao. “ISO/OSI, IEEE 802.2, and TCP/IP.”

<a href="https://sansorg.egnyte.com/dl/oyHUTHkIzn">Original paper</a>
