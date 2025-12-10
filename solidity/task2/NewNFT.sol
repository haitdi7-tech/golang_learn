// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/utils/Counters.sol";

contract NewNFT is ERC721,ERC721URIStorage{
    using Counters for Counters.Counter;

    Counters.Counter private _tokenIds;
    constructor() ERC721("NewNFT","NNFT"){}

    function mintNFT( address recipient,string memory tokenURI) public returns (uint256){
        _tokenIds.increment();

        uint256 newItemId = _tokenIds.current();
        _safeMint(recipient,newItemId);
        
        _setTokenURI(newItemId,tokenURI);

        return newItemId;
    }

    // 修复冲突：重写 tokenURI 函数
    // override(ERC721, ERC721URIStorage) 声明要覆盖的父合约
    function tokenURI(uint256 tokenId) public view override(ERC721, ERC721URIStorage) returns (string memory) {
        // 调用父类的 tokenURI 实现（优先使用 ERC721URIStorage 的逻辑）
        return super.tokenURI(tokenId);
    }

    // 修复冲突：重写 supportsInterface 函数
    function supportsInterface(bytes4 interfaceId) public view override(ERC721, ERC721URIStorage) returns (bool) {
        // 调用父类的实现，确保支持 ERC721 标准和 URI 扩展接口
        return super.supportsInterface(interfaceId);
    }
}
