What Is Unison?
===============

**Unison** is an end-to-end (E2E)/peer-to-peer (P2P) networking library for any
application that needs to self-organize an emerging network of nodes.  Unison is
built upon existing standardized technologies wherever applicable, in order to
leverage decades-long research, development, and deployment insights of the
Internet.

We at `Harmony`_ are developing and using Unison as one of the foundational
layers of our highly scalable and performant blockchain network.  We are
releasing Unison in open source because we believe that E2E/P2P networking can
enable not only blockchains but a much wider spectrum of applications.

What Does Unison Provide?
=========================

Identity-based Networking
-------------------------

Unison provides an end-to-end networking layer where hosts are *located* by
their IP addresses but *identified* logically by a cryptographic private/public
key pair that it holds.

A real-life analogue of identity-based addressing would be to send a letter
envelope-addressed to “Tim Cook” (identifier) without having to know his postal
address, and letting the post office figure out the address for you and deliver
the letter to him.  In contrast, in the traditional locator (IP address)-based
networking, a letter would be addressed to “One Infinite Loop, Cupertino, CA,
USA” (locator) without specifying the recipient name (identifier), and it would
be your responsibility to keep your address book up to date across your
correspondents' moving to new addresses.

Using Unison, a networking application on a host initiates connections to other
hosts not using their IP address *(locator)*, but their public key
*(identifier).*  Unison implements mechanisms to transparently handle the
identifier/locator mapping.

Group-based Networking
----------------------

Sometimes your application needs to send messages or objects to not just one
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

Unison automates this for your application, so that the application would not
have to concern itself with individual mechanisms and their inner workings.
would not have to concern itself with the exact mechanisms of 

Mobility
--------

Early computers had their IP address fixed, but that is no longer the case.  In
particular, mobile nodes constantly hop between different networks—home network,
coffee shop public WiFi, cellular connection—and each time they hop, their IP
address changes.

Long-running, background applications often need to continue their network
communication over such changes.  Unison provides a **mobility** mechanism that
can solve this problem for your application.

Using ``go-unison``
====================

To use ``go-unison`` in your Go application::

  $ go get simple-rules/go-unison

See the godoc **(TODO)** for details.

Under The Hood
==============

This section talks about how Unison implements each of the services mentioned
above.  **Note:** Since Unison is still under active development, these details
are subject to change.  We plan to freeze these by the time Unison reaches
version 1.0.

Host Identity Protocol Version 2 (HIPv2)
----------------------------------------

RaptorQ Fountain Code
---------------------

Licensing
=========

Copyright © 2018, Simple Rules Company.  All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this
   list of conditions and the following disclaimer.
2. Redistributions in binary form must reproduce the above copyright notice,
   this list of conditions and the following disclaimer in the documentation
   and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR
ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
(INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

The views and conclusions contained in the software and documentation are those
of the authors and should not be interpreted as representing official policies,
either expressed or implied, of the go-raptorq project.
