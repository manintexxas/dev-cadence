using System;
using MailKit.Net.Smtp;
using MailKit.Security;
using MimeKit;
using Extensions;

namespace Notifier
{
    public class SmtpConfig
    {
        public string Host {get; set;} = "smtp.gmail.com";
        public int Port {get; set;} = 587;
        public string? Sender {get; set;}
        public string? Password {get; set;}

    }
    public abstract class BaseNotifier(string message, string receiver)
    {
        public string Message { get; protected set; } = string.IsNullOrWhiteSpace(message)
                ? throw new ArgumentException("Message cannot be empty", nameof(message))
                : message;
        public string Receiver { get; protected set; } = receiver;

        public abstract void SendMessages();
    }

    public class GmailNotifier(string message, string receiver, SmtpConfig config) : BaseNotifier(message, Helper.GetValidMail(receiver))
    {
        private readonly SmtpConfig _smptconfig = config ?? throw new ArgumentNullException(nameof(config));

        public override void SendMessages()
        {
            var mimeMessage = new MimeMessage();
            mimeMessage.From.Add(new MailboxAddress("Sender", _smptconfig.Sender!));
            mimeMessage.To.Add(new MailboxAddress("You", Receiver));
            mimeMessage.Subject = "Messages Sender";

            mimeMessage.Body = new TextPart("plain")
            {
                Text = Message
            };

            using var client = new SmtpClient();

            try
            {
                client.Connect(_smptconfig.Host, _smptconfig.Port, SecureSocketOptions.StartTls);

                client.Authenticate(_smptconfig.Sender!, _smptconfig.Password!);

                client.Send(mimeMessage);

                client.Disconnect(true);
            }
            catch (Exception ex)
            {
                throw new InvalidOperationException($"Fail sending message {ex.Message}");
            }
        }

    }
}