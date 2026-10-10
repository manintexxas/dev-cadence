using CreateLog;
using System;

class Program
{
    static void Main(string[] args)
    {
        FileLogger.Configure("my.log", "");
        FileLogger.Info("Program mulai jalan");
        FileLogger.Error("Database gagal connect");
    }
}