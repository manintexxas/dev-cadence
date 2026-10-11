using System;
using System.Net.Mail;
using System.Security.Cryptography;
using CreateLog;

namespace Notifier
{
    static class GetNumber
    {
        public static int OtpRequest()
        {
            int otp = RandomNumberGenerator.GetInt32(1000, 10000);
            return otp;
        }
    }

    static class StringExtensions
    {
        public static bool IsValidEmail(string input)
        {
            if (string.IsNullOrEmpty(input))
            {
                return false;
            }

            try
            {
                var email = new MailAddress(input.Trim());
                return email.Address == input.Trim(); 
            }
            catch
            {
                return false;
            }
        }
    }

    abstract class BaseNotifier
    {
        private string? _receiver;

        public string Receiver
        {
            get => _receiver!;
            set => _receiver = value;
        }

        protected BaseNotifier(string? receiver)
        {
            if (string.IsNullOrEmpty(receiver))
            {
                throw new ArgumentException("Penerima tidak boleh kosong!", nameof(receiver));
            }

            Receiver = receiver.Trim();
        }

        protected abstract string Send();
    }
}