using Xunit;
using Notifier;
using System;

public class GmailNotifierTests
{
    [Fact]
        public void Constructor_WithValidData_ShouldCreateInstance()
    {

        var config = new SmtpConfig { Sender = "test@gmail.com", Password = "123" };
        string message = "Halo, ini pesan uji coba!";
        string receiver = "penerima@gmail.com";

        var notifier = new GmailNotifier(message, receiver, config);

        Assert.NotNull(notifier);
        Assert.Equal(message, notifier.Message);
    }

    [Fact]
    public void Constructor_WithEmptyMessage_ShouldThrowArgumentException()
    {
        var config = new SmtpConfig { Sender = "test@gmail.com", Password = "123" };
        string emptyMessage = "   "; 
        string receiver = "penerima@gmail.com";

        Assert.Throws<ArgumentException>(() => new GmailNotifier(emptyMessage, receiver, config));
    }

    [Fact]
    public void Constructor_WithNullConfig_ShouldThrowArgumentNullException()
    {
        string message = "Halo!";
        string receiver = "penerima@gmail.com";
        SmtpConfig? nullConfig = null;

        Assert.Throws<ArgumentNullException>(() => new GmailNotifier(message, receiver, nullConfig!));
    }
}