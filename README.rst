What Is ``libunison``?
======================

**``libunison``** is an end-to-end (E2E)/peer-to-peer (P2P) networking library for any
application that needs to self-organize an emerging network of nodes.  Unison is
built upon existing standardized technologies wherever applicable, in order to
leverage decades-long research, development, and deployment insights of the
Internet.

We at `Harmony`_ are developing and using Unison as one of the foundational
layers of our highly scalable and performant blockchain network.  We are
releasing Unison in open source because we believe that E2E/P2P networking can
enable not only blockchains but a much wider spectrum of applications.

What Does ``libunison`` Provide?
================================

**Note:** Since ``libunison`` is under active development, the list of features in this
section serves as the development roadmap.  We will release these features in
the order we see fit.  Please join the libunison-announce@harmony.one mailing list
for feature announcements.

Identity-based Networking
-------------------------

``libunison`` provides an end-to-end networking layer where hosts are *located* by
their IP addresses but *identified* logically by a cryptographic public key.  A
node proves its identity using the private key which matches its public-key
identifier.

A real-life analogue of identity-based addressing would be to send a letter
envelope-addressed to “Tim Cook” (identifier) without having to know his postal
address, and letting the post office figure out the address for you and deliver
the letter to him.  In contrast, in the traditional locator (IP address)-based
networking, a letter would be addressed to “One Infinite Loop, Cupertino, CA,
USA” (locator) without specifying the recipient name (identifier), and it would
be your responsibility to keep your address book up to date across your
correspondents' moving to new addresses.

Using ``libunison``, a networking application on a host initiates connections to other
hosts not using their IP address *(locator)*, but their public key
*(identifier).*  ``libunison`` implements mechanisms to transparently handle the
identifier/locator mapping.

Multicast Networking
--------------------

Sometimes an application needs to send messages or objects to not just one
other node (termed *unicast*), but to a *group* of nodes on the network.
Furthermore, your application may need to send messages/objects to *every*
member (termed *multicast*), or just *one* member (termed *anycast*), in such a
group.  (One special case of multicast is *broadcast*, where the target group
is the entire network itself.)

Unicast messaging is well supported on the global Internet, but neither
multicast nor broadcast can be easily used by end applications.  Applications
wishing to use either service needs to form and operate an *overlay network*,
which is a virtual network formed on top of the Internet.

Creating a software infrastructure for an overlay network is not an easy task
for non-networking-minded software developers.  Unison lowers this bar by
providing a robust infrastructure for overlay networks and implementing
multicast and anycast services that high-level applications can use with ease.

One interesting property of Unison's group-based networking is that such groups
are ad-hoc by default, and group identifier is in a large number space.  Anyone
may form and join any group, and because of the wide group identifier space, any
cryptographic random number generator or hash function can be used to form a
group identifier which is highly likely to be unique across the entire network,
without having to first register such an identifier with a central authority.

NAT Traversal
-------------

The 30+-year history of the Internet and its unexpected wild success has
resulted in one major problem: It has run out of IP addresses (locators) that
computers can use.  In the early days, each computer had exactly one IP address
that anyone else on the Internet can use to reach the computer.  This is no
longer the case: Multiple computers, such as on the same residental home network,
routinely share one public IP address, and the home Internet router acts as a
smart middlebox to disambiguate traffic flows that belong to different computers
behind the router.  This is called **Network Address Translation (NAT).**

One major downside of NAT is that it hinders end-to-end networking: There is no
clear way for one computer outside a home network to reach another inside the
home network, especially when the one inside has not talked to the one outside to
begin with.  Since the NAT mechanisms have evolved over the last 20 years or so,
different ways of working around this problem have surfaced.  These are called
**NAT traversal** mechanisms.  The problem is that none of these mechanisms is a
universal, one-size-fits-all approach: The application needs to dynamically
detect the situation and employ the right mechanism.

``libunison`` automates this, so that your application would not have to concern itself
with individual NAT traversal mechanisms and their inner workings.

End-to-end Data Integrity And Security
--------------------------------------

Modern applications require data protection services these days, including:

* Integrity protection: If data sent from one node becomes corrupted—maliciously
  or accidentally—before reaching its destination node, such as by a man in the
  middle (MitM) or hardware malfunction, the destination node should be able to
  detect and discard the data.
* Encryption: Data sent from one node to another node should be protected from
  prying eyes, e.g. a man in the middle should be able to glean only the fact
  that some data has been sent and its size, but not its content.

``libunison`` provides these services in end-to-end manner, similar to the guarantees
provided by TLS.  However, unlike TLS, ``libunison`` protects multiple logical flows
of data between two nodes in one shot, without having to perform expensive
security association establishment once for each logical data flow.

Mobility
--------

Early computers had their IP address fixed, but that is no longer the case.  In
particular, mobile nodes constantly hop between different networks—home network,
coffee shop public WiFi, cellular connection—and each time they hop, their IP
address changes.

Long-running, background applications often need to continue their network
communication over such changes.  Unison provides a **mobility** mechanism that
can solve this problem for your application.

Multihoming
-----------

Sometimes a networking site (a company, or more rarely, a household), for
redundancy, connects to the Internet using two or more ISPs (Internet service
providers), each of which assigns its own IP address or network to the site.
Such a site is said to be **multihomed**, where a computer consequently obtains
two or more IP addresses—one from each provider—that it can use to communicate
with the Internet.  If the Internet connectivity provided by one ISP is down,
the computer can still use the other ISP(s) to continue communication.  This
kind of communication handovers, while being possible, is typically quite
complex to set up and manage for end applications.

Unison handles such connection handovers transparently for the application so
that application communication is not disrupted as long as the Internet
connectivity is available via at least one ISP.

Cryptographic Key Lifetime Management
-------------------------------------

All cryptographic key materials should be cycled periodically, in terms of time
duration (“use this key only for 24 hours”) and/or amount of data protected by
the key (“use this key for encrypting up to 1GiB of data”), as a form of
defense against cryptanalysis.  This applies both to public/private key pairs
used as node identifiers and to symmetric keys used for data protection
services.

``libunison`` provides transparent key cycling services, so that applications do not
have to manually deal with them, and that application-level communication
persists without interruption across key cycling events.

Adversary-resistant Multicast
-----------------------------

In contrast to the Internet where directly interfacing networking entities are
routinely bound by real-life contractual obligations, ad-hoc P2P overlay
networks often include nodes that are not necessarily fully cooperative.  This
non-cooperativeness may arise out of rational, selfish, or even downright
malicious motivations.  As such, reliable communication over such P2P network
often needs to be implemented with a lower-than-Internet security assumptions,
and many P2P application protocols aim to serve, if not all nodes on the
network, at least all of the fully cooperative, “honest” nodes that conform to
the protocol, and assume the availability of a multicast mechanism that
enables a sender to send data to at least all of such honest nodes.

``libunison`` provides such a mechanism, using which a node can multicast a message
to all honest nodes, provided that the ratio of honest nodes to all nodes on
the network exceeds a minimum threshold, e.g. at least two thirds.

Cooperative, Fair-share Multicast
---------------------------------

When multicasting a large message to a large number of recipients, the 
distribution of bandwidth load placed on different nodes involved becomes an
issue.  A degenerate case of this is a technique called *manycast,* where the
sender simply transmits the same data over and over to each recipient, where
all transmission burden is placed solely on the sender.

``libunison`` provides a cooperative multicast mechanism, where the amount of data
sent and received by each node is linear to only the size of the message and
remains constant—*O*(1)–with regard to the size of the multicast group.

Stable Latency Jitters
----------------------

The Internet consists of data links that do not necessarily provide reliable
transmission of data: Packets (units of transmitted data) can become corrupt, or
even disappear during transit.  As such, traditional protocols aiming to achieve
reliable transmission of data, such as TCP, needed to incorporate mechanisms to
recover from packet losses.

Although TCP is quite robust against transient data losses, it poses one major
performance problem: Its *latency*—the amount of time a piece of data spends in
the network before being successfully delivered to the recipient—is not stable.
It fluctuates substantially around packet losses, and the magnitude of such
fluctuations, called *latency jitters,* is proportional to the nominal latency
from the sender to the recipient.  The nominal latency is quite large over
long-haul connections or over certain cellular data links (such as pre-4G), and
a proportionally large latency jitter makes it hard for real-time applications
to adopt or provide time-tight service level agreements (SLAs) or guarantees.

Unison provides a reliable data transport whose latency degrades gradually in
presence of packet losses, with much smaller latency jitters, at the expense of
a slight communication bandwidth overhead.  This applies to both unicast and
multicast.

Under The Hood
==============

This section talks about how ``libunison`` implements each of the services mentioned
above.  **Note:** Since ``libunison`` is still under active development, these details
are subject to change over time.  We plan to freeze these by the time ``libunison``
reaches version 1.0.

Host Identity Protocol Version 2
--------------------------------

Standardized in IETF `RFC 7401`_ and various companion documents, the Host
Identity Protocol Version 2 (HIPv2) suite serves as the groundwork for many of
the features provided by ``libunison``:

* Identifier–locator separation
* Cryptographic (public-key) node identifier
* End-to-end data protection services
* NAT traversal
* Mobility
* Multihoming
* Key cycling

Using HIPv2, each of the two end nodes identifying themselves with their own
public key and wishing to communicate to each other, first proves to the other
node that it indeed possesses the private key matching its own public-key
identifier, and jointly establishes a secret key using Diffie–Hellman (DH)
exchange.  Use of DH exchange ensures that the secret key is known only to the
two nodes but not to anyone else, including eavesdroppers.  This process is
known as *base exchange* in HIPv2.

The nodes then use the secret key derived from the base exchange round to
guard real application traffic, using another protocol named Encryption Security
Payload (ESP; `RFC 4303`_).  ESP provides both data integrity service
using HMAC (`RFC 2104`_) and encryption services using AES (`RFC 3602`_).

RaptorQ
-------

RaptorQ (`RFC 6330`_) is a binary object encoding/decoding scheme:

* A binary object can be encoded into a practically infinite number (2**24) of
  chunks;
* The original binary object can then be decoded from any *K* number of such
  encoded chunks with high probability, regardless their combination.  *K* is
  a constant chosen by the encoder, in the range [10..56403].
* The probability of the object recovery is: 99% with *K* chunks, 99.99% with
  *K* + 1 chunks

Insurance Against Packet Loss
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

A commonly used packet recovery mechanism employed by TCP and other protocols
involves acknowledgements and timeout-based retransmissions: After sending data
to a recipient, the sender expects a confirmation back from the recipient that
it has successfully received the data; in absence of such a confirmation within
some time, the sender assumes that the data has been lost during transit and
re-sends the same data again, hoping that the data would be delivered this time.

Using RaptorQ, repair information can be sent proactively if the sender expects
a baseline packet loss.

Cooperative Fair-Share Multicast
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

RapidChain proposes an information dispersal algorithm (IDA) which uses
Reed–Solomon code to achieve redundancy against non-cooperative nodes.  Because
Reed–Solomon code has a fixed code rate, the RapidChain IDA has a downside: It
has a fixed repair information overhead of 50% (assuming 2/3 honesty), even when
most nodes on the network are honest and little repair information is necessary.

``libunison`` uses RaptorQ instead of Reed–Solomon code, and can bring the code
rate close to the optimal rate required for actual honesty ratio observed.
Plus, the sender can generate additional repair symbols pessimistically to
recover gracefully when the honesty ratio suddenly drops below the currently
assumed rate, obviating the need to adjust the code rate and restart the round
(which would be required if a fixed-rate code such as Reed–Solomon were used).

Licensing
=========

Copyright © 2018, Simple Rules Company.  All rights reserved.

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
the Software, and to permit persons to whom the Software is furnished to do so,
subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.


.. _Harmony: https://harmony.one/
.. _RFC 7401: https://tools.ietf.org/html/rfc7401
.. _RFC 4303: https://tools.ietf.org/html/rfc4303
.. _RFC 3602: https://tools.ietf.org/html/rfc3602
.. _RFC 6330: https://tools.ietf.org/html/rfc6330
