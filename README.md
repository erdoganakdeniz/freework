
# In memory key-value store REST API

Tamamen standart kütüphaneler kullanılarak geliştirilmiş bir key-value store uygulamasıdır.Uygulama ayağa kaldırıldığında TIMESTAMP-data.gob isimli dosya oluşturulur.

Uygulama durup tekrar ayağa kalktığında, eğer
kaydedilmiş dosya varsa, tekrar varolan verileri store'a yükler.

Belirli aralıklarda (20 dakika) store'u dosyaya (TIMESTAMP-data.gob) kaydeder.

## API Kullanımı


## GET Endpoint
#### Tüm Öğeleri Getir İsteği

```
  GET /get
```
#### Tüm Öğeleri Getir Dönüş Değeri
````
{
    "code": 200,
    "method": "GET",
    "message": "All Data",
    "data": {
        "Anahtar": "Deger",
        "Anahtar2": "Deger2"
    }
}
````

#### İstenen Öğeyi Getir İsteği
```
  GET /get/key
```

| Parametre | Tip     | Açıklama                       |
| :-------- | :------- | :-------------------------------- |
| `key`      | `string` | **Gerekli**. Çağrılacak öğenin anahtar değeri |

#### İstenen Öğeyi Getir Olumlu Dönüş Değeri

````
{
    "code": 200,
    "method": "GET",
    "message": "Value Found",
    "data": "Deger"
}
````
#### İstenen Öğeyi Getir Olumsuz Dönüş Değeri

````
{
    "code": 404,
    "method": "GET",
    "message": "Value not Found",
    "data": ""
}
````

## SET Endpoint
```
  POST /set
```

| Parametre | Tip     | Açıklama                       |
| :-------- | :------- | :-------------------------------- |
| `key`      | `string` | **Gerekli**. Kaydedilecek öğenin anahtar değeri |
| `value`      | `string` | **Gerekli**. Kaydedilecek öğenin değeri |


#### Set Etmek İçin Kullanılacak JSON isteği
````
{
    "key" : "Anahtar",
    "value" : "Deger"
}
````
#### Set Ettikten Sonra Dönen Değer
````
{
    "code": 201,
    "method": "POST",
    "message": "Added into Store",
    "data": {
        "key": "Anahtar",
        "value": "Deger"
    }
}
````


## Delete Endpoint
```
  DELETE /delete/key
```

| Parametre | Tip     | Açıklama                       |
| :-------- | :------- | :-------------------------------- |
| `key`      | `string` | **Gerekli**. Silinecek öğenin anahtar değeri |


#### Silindikten Sonra Dönen Değer

````
{
    "code": 200,
    "method": "DELETE",
    "message": "Deleted into Store",
    "data": null
}
````


## Flush Endpoint
```
  GET /flush
```

| Parametre | Tip     | Açıklama                       |
| :-------- | :------- | :-------------------------------- |
| `key`      | `string` | **Gerekli**. Silinecek öğenin anahtar değeri |


#### Store Tamamen Boşaltıldıktan Sonra Dönen Değer

````
{}
````
## Dağıtım

Bu projeyi Docker ortamında çalıştırmak için ;

Build etmek için
```bash
  docker build --tag keyvaluestorerestapi .
```

Çalıştırmak için
```bash
  docker run keyvaluestorerestapi
```
  
## Lisans

[MIT](https://choosealicense.com/licenses/mit/)

  
