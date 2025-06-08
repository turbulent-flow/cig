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

    struct Customer {
        bytes32 id;
        string name;
        bytes32 idcard;
    }

    mapping(bytes32 => Customer) customers;

    function addCustomer(bytes memory encodedCustomer) public {
        Customer memory newCustomer = abi.decode(encodedCustomer, (Customer));

        require(newCustomer.id != bytes32(""), "Customer ID is invalid");
        require(
            keccak256(bytes(newCustomer.name)) != keccak256(""),
            "Customer name is invalid"
        );

        customers[newCustomer.id] = newCustomer;
    }

    function getCustomer(
        bytes32 customerID
    ) public view returns (Customer memory) {
        require(customerID != bytes32(""), "Customer ID is invalid");
        require(
            customers[customerID].id != bytes32(""),
            "Customer does not exist"
        );

        return customers[customerID];
    }

    struct Product {
        bytes32 id;
        string name;
        bytes32 brandID;
        bytes32 ownerID;
        uint releasedAt;
        int price;
    }

    mapping(bytes32 => Product) products;
    mapping(bytes32 => bytes32[]) brandProductIDs;
    mapping(bytes32 => bytes32[]) customerProductIDs;

    function addProduct(bytes memory encodedProduct) public {
        Product memory newProduct = abi.decode(encodedProduct, (Product));

        // 验证必传的参数
        require(newProduct.id != bytes32(""), "Product ID is invalid");
        require(
            keccak256(bytes(newProduct.name)) != keccak256(""),
            "Product name is invalid"
        );
        require(newProduct.brandID != bytes32(""), "Brand ID is invalid");
        require(newProduct.releasedAt != 0, "Released Date is invalid");
        require(newProduct.price != 0, "Price is invalid");

        // 验证相关 brand 的有效性
        Brand memory brand = brands[newProduct.brandID];
        require(brand.id != bytes32(""), "Brand does not exist");

        // 追加 product ID 到相关的 brand 数组
        bytes32[] storage productIDsByBrandID = brandProductIDs[
            newProduct.brandID
        ];
        productIDsByBrandID.push(newProduct.id);
        brandProductIDs[newProduct.brandID] = productIDsByBrandID;

        // 追加 product ID 到相关的 customer 数组
        bytes32[] storage productIDsByCustomerID = customerProductIDs[
            newProduct.ownerID
        ];
        productIDsByCustomerID.push(newProduct.id);
        customerProductIDs[newProduct.ownerID] = productIDsByCustomerID;

        // 追加 product
        products[newProduct.id] = newProduct;
    }

    function getProductIDsByBrandID(
        bytes32 brandID
    ) public view returns (bytes32[] memory) {
        require(brandID != bytes32(""), "Brand ID is invalid");
        require(
            brandProductIDs[brandID].length != 0,
            "The list of the product ID is empty"
        );

        return brandProductIDs[brandID];
    }

    function getProductIDsByCustomerID(
        bytes32 customerID
    ) public view returns (bytes32[] memory) {
        require(customerID != bytes32(""), "Customer ID is invalid");
        require(
            customerProductIDs[customerID].length != 0,
            "The list of the product ID is empty"
        );

        return customerProductIDs[customerID];
    }

    function getProduct(
        bytes32 productID
    ) public view returns (Product memory) {
        require(productID != bytes32(""), "Product ID is invalid");
        require(
            products[productID].id != bytes32(""),
            "Product does not exist"
        );

        return products[productID];
    }

    struct Transaction {
        bytes32 id;
        bytes32 productID;
        bytes32 fromCustomerID;
        bytes32 toCustomerID;
        int amount;
        uint transactedAt;
    }

    mapping(bytes32 => Transaction) transactions;

    function addTransaction(bytes memory encodedTransaction) public {
        Transaction memory newTransaction = abi.decode(
            encodedTransaction,
            (Transaction)
        );

        require(newTransaction.id != bytes32(""), "Transaction ID is invalid");
        require(
            newTransaction.productID != bytes32(""),
            "Product ID is invalid"
        );
        require(
            newTransaction.toCustomerID != bytes32(""),
            "The target customerID is invalid"
        );
        require(newTransaction.amount != 0, "Amount is invalid");
        require(newTransaction.transactedAt != 0, "Transacted Date is invalid");

        transactions[newTransaction.id] = newTransaction;
    }

    function getTransaction(
        bytes32 transactionID
    ) public view returns (Transaction memory) {
        require(transactionID != bytes32(""), "Transaction ID is invalid");
        require(
            transactions[transactionID].id != bytes32(""),
            "Transaction does not exist"
        );

        return transactions[transactionID];
    }
}
