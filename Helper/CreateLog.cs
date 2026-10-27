using System;
using System.IO;

namespace CreateLog
{
    public static class FileLogger
    {
        private static string? _directoryPath = null;
        private static string? _fileName = null;
        private static bool _isConfigure = false;

        private static string GetLogPath()
        {
            string dir = _directoryPath ?? Directory.GetCurrentDirectory();
            string name = _fileName ?? "app.log";
            
            // PERBAIKAN 1: Pastikan foldernya dibuat dulu jika belum ada di komputer
            if (!Directory.Exists(dir))
            {
                Directory.CreateDirectory(dir);
            }

            string logPath = Path.Combine(dir, name);

            if (!File.Exists(logPath))
            {
                File.Create(logPath).Dispose();
            }

            return logPath;
        }

        public static void Reset()
        {
            _directoryPath = null;
            _fileName = null;
            _isConfigure = false;
        }

        public static void Configure(string? fileName, string? directoryPath = null)
        {
            if (!string.IsNullOrEmpty(directoryPath) && !Directory.Exists(directoryPath))
            {
                throw new DirectoryNotFoundException("Directory not exist");
            }

            if (_isConfigure)
            {
                throw new InvalidOperationException("File already setup, cannot change!");
            }

            if (string.IsNullOrWhiteSpace(fileName))
            {
                throw new ArgumentException("File name cannot be empty");
            }

            if (!fileName.EndsWith(".txt", StringComparison.OrdinalIgnoreCase) && 
                !fileName.EndsWith(".log", StringComparison.OrdinalIgnoreCase))
            {
                throw new FormatException("File format not valid ( .txt / .log )");
            }

            _directoryPath = string.IsNullOrEmpty(directoryPath) ? Directory.GetCurrentDirectory() : directoryPath;
            _fileName = fileName;
            _isConfigure = true;
        }

        private static void Log(string level, string message)
        {
            string logEntry = $"{DateTime.Now:yyyy-MM-dd HH:mm:ss} [{level}] {message}";
            File.AppendAllText(GetLogPath(), logEntry + Environment.NewLine);
        }

        public static void Info(string message) => Log("INFO", message);
        public static void Error(string message) => Log("ERROR", message);
        public static void Warning(string message) => Log("WARNING", message);
    }
}