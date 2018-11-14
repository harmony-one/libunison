What Is ``go-unison``?
======================

``go-unison`` is an end-to-end (E2E)/peer-to-peer (P2P) networking library for
any application that needs to self-organize an emerging network of nodes.
``go-unison`` is built upon existing standardized technologies wherever
applicable in order to leverage decades-long research, development, and
deployment insights of the Internet.

We at `Harmony`_ are developing and using ``go-unison`` as one of the
foundational layers of our highly scalable and performant blockchain network.
We are releasing ``go-unison`` in open source because we believe that E2E/P2P
networking is useful not only in blockchains but in a much wider spectrum of
applications.

End-to-end Networking?
======================

**(TODO) end-to-end networking introduction in laymen's term ***

Peer-to-peer Networking?
========================

**(TODO) P2P networking introduction in laymen's term ***

What Does ``go-unison`` Provide?
================================

Identity-addressed Networking
-----------------------------

``go-unison`` provides an end-to-end networking layer where hosts are *located*
by their IP addresses but *identified* logically by a cryptographic
private/public key pair that it holds.

Using ``go-unison``, a networking application on a host initiates connections to
other hosts not using their IP address (locator), but their public key
(identifier).  ``go-unison`` implements mechanisms to transparently handle the
identifier/locator mapping.

A real-life analogue of identity-based addressing would be to send a letter
envelope-addressed to “Eugene Kim” (identifier) without having to know his
postal address, and letting the post office figure out the address for you and
deliver the letter to him.  In contrast, in the traditional locator (IP
address)-based networking, a letter would be addressed to “20852 Cherryland Dr,
Cupertino, CA 95132, USA” (locator) without specifying the recipient name
(identifier), and it would be your responsibility to keep your address book up
to date across your correspondents' moving to new addresses.

Content-addressed Networking
----------------------------

**(TODO) P2P and kademlia, aka “to whom *with* this data” addressing **

Using ``go-unison``
====================

To use ``go-unison`` in your Go application::

  $ go get simple-rules/go-unison

``go-raptorq`` has two main structs, Encoder and Decoder.  One Encoder/Decoder
object is used to encode/decode one message.  See the godoc **(TODO)** for
details.

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
