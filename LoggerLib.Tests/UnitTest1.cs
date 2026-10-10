using System;
using System.IO;
using Xunit;
using CreateLog;

namespace CreateLog.Tests
{
    // Menggunakan IDisposable berfungsi sebagai alat pembersih otomatis 
    // agar setiap kali satu tes selesai, memori langsung dikosongkan kembali.
    public class FileLoggerTests : IDisposable
    {
        public FileLoggerTests()
        {
            // Dibersihkan SEBELUM tes dijalankan
            FileLogger.Reset();
        }

        public void Dispose()
        {
            // Dibersihkan SETELAH tes selesai
            FileLogger.Reset();
        }

        [Fact]
        public void Configure_WithInvalidFileFormat_ThrowsFormatException()
        {
            // SKENARIO FORMAT SALAH:
            // Menguji apakah program menolak jika user memasukkan ekstensi selain .txt atau .log (misal .pdf).
            // Ini penting supaya file log tidak sembarangan formatnya.
            
            Assert.Throws<FormatException>(() => 
            {
                FileLogger.Configure("laporan.pdf");
            });
        }

        [Fact]
        public void Configure_WithEmptyFileName_ThrowsArgumentException()
        {
            // SKENARIO NAMA KOSONG:
            // Menguji apakah program menolak jika nama file dikosongkan atau hanya berisi spasi.
            
            Assert.Throws<ArgumentException>(() => 
            {
                FileLogger.Configure("   ");
            });
        }

        [Fact]
        public void Configure_CalledTwice_ThrowsInvalidOperationException()
        {
            // SKENARIO KONFIGURASI GANDA:
            // Menguji aturan bahwa logger tidak boleh di-setup dua kali secara beruntun.
            
            // Panggilan pertama (Harusnya Berhasil)
            FileLogger.Configure("app1.log");

            // Panggilan kedua (Harusnya Ditolak dan mengeluarkan error InvalidOperationException)
            Assert.Throws<InvalidOperationException>(() => 
            {
                FileLogger.Configure("app2.log");
            });
        }
    }
}