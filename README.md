  <a href="https://golang.org/" target="_blank" rel="noreferrer"> 
        <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/go/go-original.svg" alt="golang" width="40" height="40"/> 
    </a>
Encryption Text <br>
Araç, AES (Advanced Encryption Standard) algoritmasını kullanarak metinleri şifreler ve şifreli metinleri çözerek orijinal metni elde eder.
2. Fonksiyonlar ve Görevleri
2.1. generateRandomBytes(size int) ([]byte, error)

Bu fonksiyon, belirli bir boyutta rastgele bayt dizisi (byte array) oluşturur. Kriptografik olarak güvenli rastgele baytlar üretmek için kullanılır.
2.2. encrypt(plainText, key string) (string, error)

Bu fonksiyon, verilen düz metni (plain text) ve anahtarı (key) kullanarak şifreler. Şifreleme işlemi AES algoritmasının CBC (Cipher Block Chaining) modu kullanılarak gerçekleştirilir.

    Girdi: Düz metin ve şifreleme anahtarı.
    Çıktı: Şifrelenmiş metin (hex formatında).

Şifreleme süreci şu adımları içerir:

    Düz metin ve anahtar bayt dizisine dönüştürülür.
    AES blok şifreleyici oluşturulur.
    Rastgele bir IV (Initialization Vector) oluşturulur.
    Düz metin, AES blok boyutuna göre doldurulur (padding).
    CBC modunda şifreleme yapılır.
    IV ve şifrelenmiş metin birleştirilerek sonuç hex formatında döndürülür.

2.3. decrypt(cipherText, key string) (string, error)

Bu fonksiyon, verilen şifreli metni ve anahtarı kullanarak şifresini çözer. Şifre çözme işlemi AES algoritmasının CBC modu kullanılarak gerçekleştirilir.

    Girdi: Şifreli metin (hex formatında) ve şifre çözme anahtarı.
    Çıktı: Orijinal düz metin.

Şifre çözme süreci şu adımları içerir:

    Şifreli metin hex formatından bayt dizisine dönüştürülür.
    AES blok şifreleyici oluşturulur.
    IV, şifreli metinden ayrılır.
    CBC modunda şifre çözme yapılır.
    Düz metin, doldurma baytlarından arındırılır (unpadding).
    Orijinal düz metin döndürülür.

2.4. main()

Bu fonksiyon, anahtar oluşturma, metin şifreleme ve şifre çözme işlemlerini gerçekleştiren temel akış kontrolünü içerir.

    Girdi: Sabit bir anahtar ve düz metin.
    Çıktı: Şifrelenmiş metin ve şifresi çözülmüş metin.

3. Çıktı

Program çalıştırıldığında aşağıdaki çıktıları üretir:

    Şifrelenmiş metin (hex formatında).
    Şifresi çözülmüş metin (orijinal düz metin).

4. Sonuç

Bu metin şifreleme ve şifre çözme aracı, AES algoritmasını kullanarak güvenli bir şekilde metin şifrelemek ve şifre çözmek için kullanılabilir. AES'in sağlam şifreleme yetenekleri sayesinde, kullanıcıların verileri güvenli bir şekilde koruma altına alınabilir. Programın doğru çalışabilmesi için, kullanılan anahtarın doğru ve yeterli uzunlukta olması gerekmektedir.










