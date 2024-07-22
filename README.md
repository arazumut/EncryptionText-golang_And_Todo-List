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

 Todo List <a href="https://golang.org/" target="_blank" rel="noreferrer"> 
        <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/go/go-original.svg" alt="golang" width="40" height="40"/> 
    </a> <br>
    Bu Go programı, kullanıcıların komut satırında etkileşimli olarak bir yapılacaklar listesi (to-do list) yönetmesine olanak tanır. Program, görevleri ekleme, silme ve tamamlama işlevlerini sunar ve görevleri JSON formatında bir dosyada saklar. Aşağıda programın nasıl çalıştığını ve her bir bölümünün ne yaptığını açıklıyorum.

Programın Genel Yapısı
Görevlerin Tanımlanması: Görevler (tasks), her biri bir kimlik (ID), isim (name) ve tamamlanma durumu (done) olan Task adında bir yapı (struct) ile temsil edilir.
Veri Depolama: Görevler, tasks.json adında bir dosyada JSON formatında saklanır. Görevler program başlatıldığında bu dosyadan yüklenir ve program sonlandırıldığında bu dosyaya kaydedilir.
Kullanıcı Etkileşimi: Kullanıcı, komut satırında etkileşimli olarak görev ekleyebilir, silebilir ve tamamlanmış olarak işaretleyebilir. Kullanıcıdan alınan komutlara göre uygun işlemler gerçekleştirilir.
Detaylı Açıklama
main Fonksiyonu
Program başlatıldığında loadTasks fonksiyonu çağrılarak görevler tasks.json dosyasından yüklenir.
Sonsuz bir döngü içinde kullanıcıdan komutlar alınır ve uygun işlemler gerçekleştirilir. Kullanıcı quit komutunu girene kadar bu döngü devam eder.
loadTasks Fonksiyonu
tasks.json dosyasını okur ve JSON formatındaki veriyi tasks dilim (slice) değişkenine dönüştürür.
Dosya bulunamazsa (ilk çalıştırma durumu) boş bir görev listesi oluşturur.
Dosya okuma veya JSON çözümleme sırasında hata oluşursa program sonlandırılır.
saveTasks Fonksiyonu
tasks dilim değişkenini JSON formatına dönüştürür ve tasks.json dosyasına yazar.
JSON dönüştürme veya dosya yazma sırasında hata oluşursa program sonlandırılır.
listTasks Fonksiyonu
tasks dilimindeki görevleri listeler. Görevler arasında tamamlanmış olanları "x" ile işaretler.
Eğer görev listesi boşsa "No tasks" mesajı görüntülenir.
addTask Fonksiyonu
Kullanıcıdan bir görev adı alır.
Yeni bir görev oluşturur ve tasks dilimine ekler. Yeni görevin ID'si, mevcut görevlerin en büyük ID'sinden bir fazladır.
Yeni görev ekledikten sonra görev listesini dosyaya kaydeder.
removeTask Fonksiyonu
Kullanıcıdan bir görev ID'si alır.
Verilen ID'ye sahip görevi tasks diliminden kaldırır.
Görev silindikten sonra görev listesini dosyaya kaydeder.
Görev bulunamazsa hata mesajı görüntülenir.
markTaskAsDone Fonksiyonu
Kullanıcıdan bir görev ID'si alır.
Verilen ID'ye sahip görevi done olarak işaretler.
Görev tamamlandıktan sonra görev listesini dosyaya kaydeder.
Görev bulunamazsa hata mesajı görüntülenir.
Program Akışı
Program başlatılır.
Görevler tasks.json dosyasından yüklenir.
Kullanıcıya komut girmesi için bir menü sunulur.
Kullanıcı bir komut girer:
add: Yeni görev ekler.
remove: Mevcut bir görevi siler.
done: Mevcut bir görevi tamamlanmış olarak işaretler.
quit: Programdan çıkar ve görevleri dosyaya kaydeder.
Kullanıcı quit komutunu girene kadar adım 3 ve 4 tekrar eder.
Sonuç
Bu Go programı, kullanıcıların bir yapılacaklar listesi yönetmesine olanak tanır. Görevler eklenebilir, silinebilir ve tamamlanmış olarak işaretlenebilir. Tüm görevler JSON formatında bir dosyada saklanır ve program her başlatıldığında bu dosyadan yüklenir. Program, komut satırında etkileşimli olarak çalışır ve kullanıcıdan alınan komutlara göre uygun işlemleri gerçekleştirir.
















