package handler

import (
    "SecretCare/entity"
    "testing"
    "github.com/stretchr/testify/assert"
    "fmt"
)

// TestCreateNewProduct menguji proses fungsi CreateNewProduct
func TestCreateNewProduct(t *testing.T) {
    // Inisialisasi mock repository
    productMock := &ProductMock{}

    // Tentukan produk baru yang akan dibuat
    newProduct := entity.Product{
        ID:     1,
        Nama:   "Keyboard",
        Harga:  200,
        Stock:  20,
        TokoID: 3,
    }

    // Tentukan perilaku mock
    productMock.On("CreateNewProduct", newProduct).Return(nil)

    // Buat handler dan inject mock
    handler := HandlerProduct(productMock)

    //  Panggil metode CreateNewProduct
    err := handler.CreateNewProduct(newProduct)
    
    // Verifikasi bahwa error tidak ada
    assert.Nil(t, err)

    // Verifikasi apakah metode mock dipanggil dengan benar
    productMock.AssertCalled(t, "CreateNewProduct", newProduct)

    // Pastikan semua ekspektasi mock dipenuhi
    productMock.AssertExpectations(t)
}

// TestCreateNewProduct_Failure menguji kegagalan saat pembuatan produk.
func TestCreateNewProduct_Failure(t *testing.T) {
    // Inisialisasi mock repository
    productMock := &ProductMock{}

    // Step 2: Tentukan produk baru yang akan dibuat
    newProduct := entity.Product{
        ID:     2,
        Nama:   "Monitor",
        Harga:  150,
        Stock:  10,
        TokoID: 4,
    }

    // Tentukan perilaku mock untuk simulasi error
    productMock.On("CreateNewProduct", newProduct).Return(fmt.Errorf("database error"))

    // Buat handler dan inject mock
    handler := HandlerProduct(productMock)

    // Panggil metode CreateNewProduct
    err := handler.CreateNewProduct(newProduct)

    // Verifikasi bahwa error dikembalikan
    assert.NotNil(t, err)
    assert.EqualError(t, err, "database error")

    // Pastikan ekspektasi mock dipenuhi
    productMock.AssertExpectations(t)
}

// TestCreateNewProduct_InvalidData menguji pembuatan produk dengan data yang tidak valid.
func TestCreateNewProduct_InvalidData(t *testing.T) {
    // Inisialisasi mock repository
    productMock := &ProductMock{}

    // Tentukan produk yang tidak valid (harga negatif)
    invalidProduct := entity.Product{
        ID:     3,
        Nama:   "Headset",
        Harga:  -50, // Harga tidak valid
        Stock:  10,
        TokoID: 5,
    }

    // Tentukan perilaku mock yang mengembalikan error
    productMock.On("CreateNewProduct", invalidProduct).Return(fmt.Errorf("invalid product data"))

    // Buat handler menggunakan mock
    handler := HandlerProduct(productMock)

    // Panggil metode CreateNewProduct
    err := handler.CreateNewProduct(invalidProduct)

    // Verifikasi bahwa error dikembalikan
    assert.NotNil(t, err)
    assert.Contains(t, err.Error(), "invalid product data")

    // Pastikan metode mock dipanggil dengan benar
    productMock.AssertCalled(t, "CreateNewProduct", invalidProduct)

    // Pastikan semua ekspektasi mock dipenuhi
    productMock.AssertExpectations(t)
}


// TestUpdateStockById_Success menguji pembaruan stok produk yang berhasil.
func TestUpdateStockById_Success(t *testing.T) {
    // Inisialisasi mock repository
    productMock := &ProductMock{}

    // Tentukan ID produk dan stok baru
    productID := 1
    newStock := 50

    // Tentukan perilaku mock untuk pembaruan yang sukses
    productMock.On("UpdateStockById", productID, newStock).Return(nil)

    // Gunakan mock repository sebagai handler
    handler := HandlerProduct(productMock)

    // Panggil metode UpdateStockById
    err := handler.UpdateStockById(productID, newStock)

    // Verifikasi bahwa tidak ada error
    assert.Nil(t, err)

    // Step 7: Verifikasi bahwa mock dipanggil dengan argumen yang benar
    productMock.AssertCalled(t, "UpdateStockById", productID, newStock)
    productMock.AssertExpectations(t)
}

// TestUpdateStockById_Error menguji pembaruan stok produk yang gagal.
func TestUpdateStockById_Error(t *testing.T) {
    // Inisialisasi mock repository
    productMock := &ProductMock{}

    // Tentukan ID produk dan stok baru yang tidak valid
    productID := 1
    newStock := -10 // Nilai stok negatif

    // Tentukan perilaku mock untuk error
    productMock.On("UpdateStockById", productID, newStock).Return(fmt.Errorf("invalid stock value"))

    // Gunakan mock repository sebagai handler
    handler := HandlerProduct(productMock)

    // Panggil metode UpdateStockById
    err := handler.UpdateStockById(productID, newStock)

    // Verifikasi bahwa error terjadi
    assert.NotNil(t, err)
    assert.Contains(t, err.Error(), "invalid stock value")

    // Verifikasi bahwa mock dipanggil dengan argumen yang benar
    productMock.AssertCalled(t, "UpdateStockById", productID, newStock)
    productMock.AssertExpectations(t)
}

// TestDeleteProductById_Success menguji penghapusan produk yang berhasil.
func TestDeleteProductById_Success(t *testing.T) {
    // Inisialisasi mock repository
    productMock := &ProductMock{}

    // Tentukan perilaku mock untuk penghapusan yang berhasil
    productMock.On("DeleteProductById", 1).Return(nil)

    // Panggil metode DeleteProductById
    err := productMock.DeleteProductById(1)

    // Verifikasi bahwa tidak ada error
    assert.Nil(t, err)

    // Verifikasi ekspektasi
    productMock.AssertCalled(t, "DeleteProductById", 1)
    productMock.AssertExpectations(t)
}

// TestDeleteProductById_DatabaseError menguji penghapusan produk yang gagal karena error di database.
func TestDeleteProductById_DatabaseError(t *testing.T) {
    // Inisialisasi mock repository
    productMock := &ProductMock{}

    // Simulasikan error di database (misalnya kesalahan saat melakukan query)
    productMock.On("DeleteProductById", 1).Return(fmt.Errorf("database error"))

    // Panggil metode DeleteProductById
    err := productMock.DeleteProductById(1)

    // Verifikasi bahwa error terjadi
    assert.NotNil(t, err)
    assert.Contains(t, err.Error(), "database error")

    // Verifikasi ekspektasi
    productMock.AssertCalled(t, "DeleteProductById", 1)
    productMock.AssertExpectations(t)
}

// TestGetProductsByTokoID menguji pengambilan produk berdasarkan ID toko.
func TestGetProductsByTokoID(t *testing.T) {
    // Inisialisasi mock repository
    productMock := &ProductMock{}

    // Tentukan hasil yang diharapkan
    expectedProducts := []entity.Product{
        {ID: 1, Nama: "Laptop", Harga: 1000, Stock: 10, TokoID: 2},
        {ID: 2, Nama: "Mouse", Harga: 50, Stock: 30, TokoID: 2},
    }

    // Tentukan perilaku mock untuk GetProductsByTokoID dengan tokoID = 2
    productMock.On("GetProductsByTokoID", 2).Return(expectedProducts, nil)

    // Gunakan mock repository untuk pengujian
    handlerInterface := HandlerProduct(productMock)

    // Panggil metode GetProductsByTokoID
    result := handlerInterface.GetProductsByTokoID(2)

    // Verifikasi bahwa hasilnya sesuai dengan yang diharapkan
    assert.Equal(t, expectedProducts, result)

    // Verifikasi ekspektasi mock
    productMock.AssertExpectations(t)
}

// TestGetProductsByTokoIDNotFound menguji pengambilan produk dengan tokoID yang tidak ada produk.
func TestGetProductsByTokoIDNotFound(t *testing.T) {
    // Inisialisasi mock repository
    productMock := &ProductMock{}

    // Tentukan hasil yang diharapkan (tidak ada produk ditemukan)
    emptyProducts := []entity.Product{}

    // Tentukan perilaku mock untuk GetProductsByTokoID dengan tokoID = 10
    productMock.On("GetProductsByTokoID", 10).Return(emptyProducts, nil)

    // Gunakan mock repository untuk pengujian
    handlerInterface := HandlerProduct(productMock)

    // Panggil metode GetProductsByTokoID
    result := handlerInterface.GetProductsByTokoID(10)

    // Verifikasi bahwa tidak ada produk ditemukan
    assert.Empty(t, result)

    // Verifikasi ekspektasi mock
    productMock.AssertExpectations(t)
}

// TestGetProductReport_Success menguji apakah laporan produk berhasil diambil dengan benar
// untuk toko yang memiliki tokoID yang valid (misalnya 2). Di sini, kita mock repository
// untuk mengembalikan data laporan produk yang sudah dipersiapkan sebelumnya. Fungsi ini
// memastikan bahwa hasil yang dikembalikan oleh GetProductReport sesuai dengan data yang diharapkan.
func TestGetProductReport_Success(t *testing.T) {
    // Inisialisasi mock repository
    productMock := &ProductMock{}

    // Menyiapkan data laporan produk yang diharapkan
    expectedReport := []entity.ProductReport{
        {Nama: "Pelembab", TotalPenjualan: 20, TotalPendapatan: 200000},
        {Nama: "Toner", TotalPenjualan: 10, TotalPendapatan: 100000},
    }

    // Menyiapkan perilaku mock: ketika GetProductReport dipanggil dengan tokoID = 2,
    // kembalikan data laporan produk yang sudah disiapkan
    productMock.On("GetProductReport", 2).Return(expectedReport)

    // Panggil metode GetProductReport dengan tokoID = 2
    result := productMock.GetProductReport(2)

    // Pastikan hasilnya sesuai dengan laporan produk yang diharapkan
    assert.Equal(t, expectedReport, result)

    // Verifikasi bahwa mock repository telah dipanggil dengan parameter yang benar
    productMock.AssertCalled(t, "GetProductReport", 2)
    productMock.AssertExpectations(t)
}

// TestGetProductReport_NoProductsFound menguji skenario di mana tidak ada produk yang ditemukan
// untuk toko dengan tokoID tertentu. Dalam hal ini, kita mock repository untuk mengembalikan
// laporan kosong, yang berarti tidak ada produk untuk tokoID yang diberikan.
func TestGetProductReport_NoProductsFound(t *testing.T) {
    // Inisialisasi mock repository
    productMock := &ProductMock{}

    // Menyiapkan data laporan produk kosong, artinya tidak ada produk untuk tokoID ini
    var emptyReport []entity.ProductReport

    // Menyiapkan perilaku mock: ketika GetProductReport dipanggil dengan tokoID = 5,
    // kembalikan laporan kosong
    productMock.On("GetProductReport", 5).Return(emptyReport)

    // Panggil metode GetProductReport dengan tokoID = 5
    result := productMock.GetProductReport(5)

    // Pastikan hasilnya adalah slice kosong (tidak ada produk ditemukan)
    assert.Empty(t, result)

    // Verifikasi bahwa mock repository telah dipanggil dengan parameter yang benar
    productMock.AssertCalled(t, "GetProductReport", 5)
    productMock.AssertExpectations(t)
}