using System;
using MailKit.Net.Smtp;
using MailKit.Security;
using MimeKit;
using Extensions;
using System.IO;
using Microsoft.VisualStudio.TestPlatform.CommunicationUtilities;
using Microsoft.AspNetCore.SignalR;
namespace Notifier
{
    public class SmtpConfig
    {
        public string Host { get; set; } = "smtp.gmail.com";
        public int Port { get; set; } = 587;
        public string? Sender { get; set; }
        public string? Password { get; set; }

    }
    public abstract class BaseNotifier(string receiver)
    {
        public string? Message { get; protected set; }
        public string Receiver { get; protected set; } = receiver;

        public abstract void SendMessages();
    }

    public class GmailNotifier : BaseNotifier
    {
        public SmtpConfig? smtpConfig {get; private set;} 
        public string SecretNumber {get; private set;} = Helper.GetNumber();

        public GmailNotifier(string receiver, SmtpConfig config) : base(Helper.GetValidMail(receiver, "gmail.com"))
        {
            string path = @"C:\Users\LENOVO\Documents\Snowball\My-labs\P01-Notifier\Template.html";

            string Template = File.ReadAllText(path).Replace("{DATE_TIME}", DateTime.Now.ToString("yyyy-MM-dd"));

            Message = Template;

            smtpConfig = config ?? throw new ArgumentNullException(nameof(config));
        }

        public override void SendMessages()
        {
            var mimeMessage = new MimeMessage();
            mimeMessage.From.Add(new MailboxAddress("Sender", smtpConfig!.Sender!));
            mimeMessage.To.Add(new MailboxAddress("You", Receiver));
            mimeMessage.Subject = "Messages Sender";

            mimeMessage.Body = new TextPart(MimeKit.Text.TextFormat.Html)
            {
                Text = Message!
            };

            using var client = new SmtpClient();

            try
            {
                client.Connect(smtpConfig.Host, smtpConfig.Port, SecureSocketOptions.StartTls);

                client.Authenticate(smtpConfig.Sender!, smtpConfig.Password!);

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