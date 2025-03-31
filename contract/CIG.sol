// SPDX-License-Identifier: MIT

pragma solidity ^0.8.26;

contract CIG {
    struct Brand {
        bytes32 id;
        string name;
        string description;
        uint createdAt;
    }

    mapping(bytes32 => Brand) brands;
    Brand[] brandIDs;

    function addBrand(bytes memory encodedBrand) public {
        Brand memory newBrand = abi.decode(encodedBrand, (Brand));

        require(newBrand.id != bytes32(""), "Brand ID is invalid");
        require(
            keccak256(bytes(newBrand.name)) != keccak256(""),
            "Brand name is invalid"
        );

        brandIDs.push();

        brands[newBrand.id] = newBrand;
    }

    function getBrand(bytes32 brandID) public view returns (Brand memory) {
        require(brandID != bytes32(""), "Brand ID is invalid");
        require(brands[brandID].id != bytes32(""), "Brand does not exist");

        return brands[brandID];
    }
}
