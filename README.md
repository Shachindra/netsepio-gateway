# README #

Gateway for Minting Reviews as NFTs on Aptos Blockchain

aptos move run --function-id '0xc333017dfac8488f3339d39c56b0b33191af5985d1335eeefa0b3bf59fc9f4a9::netsepio::delegate_submit_review' \
--args 'address:0xe6ff9904344bad934dd708233fcfb79aa385cd93fc424994ae0ed97fb71160c6' \
--args 'string:ipfs://ipfshash' \
--args 'string:category' \
--args 'string:mydomain.com' \
--args 'string:full/url/link' \
--args 'string:type' \
--args 'string:tag1, tag2, tag3, tag4' \
--args 'string:safe' \
--url "https://fullnode.devnet.aptoslabs.com" --private-key "0xc596dd3792261400c82ba1061c2a46b1dc1588143d3653d576b77192adf2bc91" --assume-yes

aptos move run --function-id '0xc333017dfac8488f3339d39c56b0b33191af5985d1335eeefa0b3bf59fc9f4a9::netsepio::delegate_submit_review' --args 'address:0xe6ff9904344bad934dd708233fcfb79aa385cd93fc424994ae0ed97fb71160c6' --args 'string:ipfs://ipfshash2' --args 'string:category' --args 'string:mydomain.com' --args 'string:full/url/link' --args 'string:type' --args 'string:tag1, tag2, tag3, tag4' --args 'string:safe' --url "https://fullnode.devnet.aptoslabs.com" --private-key "0x5d062b4e6fa5e47e114d076801dfd37e57ebbab81de4d9a5c00dda57b7c8395b" --assume-yes