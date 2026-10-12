using System;
using MimeKit;  
using System.Security.Cryptography;
using MailKit.Net.Smtp;
using MailKit.Security;
using DotNetEnv;

namespace Extensions
{
    static class Helper
    {
        public static string GetValidMail(string? input, string? host = null)
        {
            string getHost = host ?? "gmail.com";

            if (string.IsNullOrWhiteSpace(input))
            {
                throw new ArgumentNullException(nameof(input), "Email cannot be empty");
            }

            if (!MailboxAddress.TryParse(input.Trim(), out var mailbox))
            {
                throw new FormatException($"{input} format is invalid");
            }

            if (!string.Equals(mailbox.Domain, getHost, StringComparison.OrdinalIgnoreCase))
            {
                throw new ArgumentException($"{input} domain name is ivalid");
            }

            return mailbox.Address.ToLower();
        }

        public static string GetNumber()
        {
            return RandomNumberGenerator.GetInt32(0, 10000).ToString("D4");
        }

        public static void GetEnv()
        {
            var currentDirectory = new DirectoryInfo(Directory.GetCurrentDirectory());

            while (currentDirectory != null)
            {
                string envPath = Path.Combine(currentDirectory.FullName, ".env");
                if (File.Exists(envPath))
                {
                    Env.Load(envPath);
                    return;
                }

                currentDirectory = currentDirectory.Parent;
            }
        }
    }
}