/*
 *  Copyright 2019 The searKing authors. All Rights Reserved.
 *
 *  Use of this source code is governed by a MIT-style license
 *  that can be found in the LICENSE file in the root of the source
 *  tree. An additional intellectual property rights grant can be found
 *  in the file PATENTS.  All contributing project authors may
 *  be found in the AUTHORS file in the root of the source tree.
 */
#ifndef GO_OS_SIGNAL_CGO_SIGNAL_HANDLER_H_
#define GO_OS_SIGNAL_CGO_SIGNAL_HANDLER_H_
#include <signal.h>
#include <stdio.h>
#include <unistd.h>

#include <atomic>
#include <functional>
#include <map>
#include <mutex>
#include <utility>

namespace searking {

// Callbacks Predefinations
typedef void (*SIGNAL_SA_ACTION_CALLBACK)(int signum, siginfo_t *info,
                                          void *context);
typedef void (*SIGNAL_SA_HANDLER_CALLBACK)(int signum);

typedef void (*SIGNAL_ON_SIGNAL_CALLBACK)(void *ctx, int fd, int signum,
                                          siginfo_t *info, void *context);

typedef void (*SIGNAL_ON_BACKTRACE_DUMP_CALLBACK)(int fd);

class SignalHandler {
 protected:
  SignalHandler()
      : onSignalCtx_(nullptr), onSignal_(nullptr), fd_(fileno(stdout)) {}

 public:
  // Thread safe GetInstance.
  static SignalHandler &GetInstance();

  void operator()(int signum, siginfo_t *info, void *context);
  void RegisterOnSignal(std::function<void(void *ctx, int fd, int signum,
                                           siginfo_t *info, void *context)>
                            callback,
                        void *ctx);
  void SetBacktraceDumpTo(std::function<void(int fd)> safe_dump_to);

  void SetFd(int fd);

  void SetFd(FILE *fd);

  void SetSigactionHandlers(int signum, SIGNAL_SA_ACTION_CALLBACK action,
                            SIGNAL_SA_HANDLER_CALLBACK handler);

  static int SignalAction(int signum);
  static int SignalAction(int signum, SIGNAL_SA_ACTION_CALLBACK action,
                          SIGNAL_SA_HANDLER_CALLBACK handler);

 private:
  std::mutex mutex_;
  int fd_;
  std::function<void(int fd)> backtrace_dump_to_;
  void *onSignalCtx_;
  std::function<void(void *ctx, int fd, int signum, siginfo_t *info,
                     void *context)>
      onSignal_;
  std::map<int,
           std::pair<SIGNAL_SA_ACTION_CALLBACK, SIGNAL_SA_HANDLER_CALLBACK> >
      sigactionHandlers_;

 private:
  SignalHandler(const SignalHandler &) = delete;
  void operator=(const SignalHandler &) = delete;
};
}  // namespace searking
#endif  // GO_OS_SIGNAL_CGO_SIGNAL_HANDLER_H_
