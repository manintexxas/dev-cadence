using Notifier;
using Extensions;

var builder = WebApplication.CreateBuilder(args);
Helper.GetEnv();

builder.Services.AddSingleton(new SmtpConfig
{
    Sender = Environment.GetEnvironmentVariable("SMTP_USER"),
    Password = Environment.GetEnvironmentVariable("SMTP_PASS"),
    Host = Environment.GetEnvironmentVariable("SMTP_HOST") ?? "smtp.gmail.com",
    Port = int.Parse(Environment.GetEnvironmentVariable("SMTP_PORT") ?? "587")
});

var app = builder.Build();
app.UseDefaultFiles(); 
app.UseStaticFiles();

app.MapPost("/api/email/send", (EmailRequestDto request, SmtpConfig config) =>
{
    try
    {
        GmailNotifier notifier = new GmailNotifier(request.Receiver, config);

        notifier.SendMessages();

        return Results.Ok(new { status = "success", message = "Email berhasil dikirim!" });
    }
    catch (Exception ex)
    {
        return Results.BadRequest(new { status = "error", message = ex.Message });
    }
});

app.Run();